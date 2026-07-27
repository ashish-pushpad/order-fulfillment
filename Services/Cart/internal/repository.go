package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	// "time"
)

var (
	ErrCartNotFound = errors.New("cart not found")
	ErrItemNotFound = errors.New("cart item not found")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetOrCreateCart(ctx context.Context, userID int64) (int64, error) {
	var cartID int64

	query := `
		INSERT INTO carts (user_id)
		VALUES ($1)
		ON CONFLICT (user_id) DO UPDATE 
		SET user_id = EXCLUDED.user_id 
		RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&cartID)
	if err != nil {
		return 0, fmt.Errorf("failed to get or create cart: %w", err)
	}

	return cartID, nil
}

func (r *Repository) AddItem(ctx context.Context, cartID int64, productID int64, qty int,price float32) error {
	
	query := `
		INSERT INTO cart_items (cart_id, product_id, quantity, price_when_added)
		VALUES (
			$1, 
			$2, 
			$3,
			$4
		)
		ON CONFLICT (cart_id, product_id) DO UPDATE 
		SET quantity = cart_items.quantity + EXCLUDED.quantity
	`
	_, err := r.db.ExecContext(ctx, query, cartID, productID, qty,price)
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}

	return nil
}

func (r *Repository) RemoveItem(ctx context.Context, cartItemID int64) error {

	query := `DELETE FROM cart_items WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, cartItemID)
	if err != nil {
		return fmt.Errorf("failed to remove item: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrItemNotFound
	}

	return nil
}

func (r *Repository) UpdateQuantity(ctx context.Context, cartItemID int64, qty int) error {
	query := `
		UPDATE cart_items
		SET quantity = $1
		WHERE id = $2
	`
	result, err := r.db.ExecContext(ctx, query, qty, cartItemID)
	if err != nil {
		return fmt.Errorf("failed to update quantity: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrItemNotFound
	}

	return nil
}

func (r *Repository) ClearCartTx(
	ctx context.Context,
	tx *sql.Tx,
	cartID int64,
) error {

	_, err := tx.ExecContext(
		ctx,
		`DELETE FROM cart_items WHERE cart_id = $1`,
		cartID,
	)

	if err != nil {
		return fmt.Errorf("failed to clear cart: %w", err)
	}

	return nil
}

func (r *Repository) ClearCart(
	cartID int64,
) error {

	_, err := r.db.Exec(
		`DELETE FROM cart_items WHERE cart_id = $1`,
		cartID,
	)

	if err != nil {
		return fmt.Errorf("failed to clear cart: %w", err)
	}

	return nil
}
// have to use the getProduct from the product grpcServer intead of the join
func (r *Repository) GetCart(ctx context.Context, userID int64) (CartResponse, error) {
	var cart CartResponse

	err := r.db.QueryRowContext(ctx, `
		SELECT id
		FROM carts
		WHERE user_id = $1
	`, userID).Scan(&cart.ID)

	if errors.Is(err, sql.ErrNoRows) {
		return CartResponse{}, ErrCartNotFound
	}
	if err != nil {
		return CartResponse{}, fmt.Errorf("failed to fetch cart: %w", err)
	}

	cart.UserID = userID

	rows, err := r.db.QueryContext(ctx, `
		SELECT
			id,
			quantity,
			price_when_added,
			product_id
		FROM cart_items 
		WHERE cart_id = $1
		ORDER BY id
	`, cart.ID)
	if err != nil {
		return CartResponse{}, fmt.Errorf("failed to fetch cart items: %w", err)
	}
	defer rows.Close()

	// Prefill slice with non-nil empty slice to avoid returning `null` in JSON responses
	cart.Items = make([]CartItemResponse, 0)

	for rows.Next() {
		var item CartItemResponse
		err := rows.Scan(
			&item.ID,
			&item.Quantity,
			&item.UnitPrice,
			&item.ProductID,
		)
		if err != nil {
			return CartResponse{}, fmt.Errorf("failed to scan item: %w", err)
		}

		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		cart.Subtotal += item.TotalPrice
		cart.TotalItems += item.Quantity

		cart.Items = append(cart.Items, item)
	}

	if err := rows.Err(); err != nil {
		return CartResponse{}, fmt.Errorf("rows iteration error: %w", err)
	}

	return cart, nil
}

// GetExpiration retrieves the expiration window for a given cart.
// func (r *Repository) GetExpiration(ctx context.Context, cartID int64) (*CartTime, error) {
// 	var ct CartTime
// 	query := `SELECT cart_id, expires_at FROM cart_times WHERE cart_id = $1`
// 	err := r.db.QueryRowContext(ctx, query, cartID).Scan(&ct.CartID, &ct.ExpiresAt)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, ErrCartNotFound
// 		}
// 		return nil, err
// 	}
// 	return &ct, nil
// }

// UpdateExpiration sets or extends the expiration timestamp for a cart (UPSERT).
// func (r *Repository) UpdateExpiration(ctx context.Context, cartID int64, ttl time.Duration) error {
// 	expirationTime := time.Now().Add(ttl)
// 	query := `
// 		INSERT INTO cart_times (cart_id, expires_at)
// 		VALUES ($1, $2)
// 		ON CONFLICT (cart_id) DO UPDATE
// 		SET expires_at = EXCLUDED.expires_at
// 	`
// 	_, err := r.db.ExecContext(ctx, query, cartID, expirationTime)
// 	return err
// }

// DeleteExpiration cleans up the TTL row if a cart is emptied or expired.
// func (r *Repository) DeleteExpiration(ctx context.Context, cartID int64) error {
// 	query := `DELETE FROM cart_times WHERE cart_id = $1`
// 	_, err := r.db.ExecContext(ctx, query, cartID)
// 	return err
// }
