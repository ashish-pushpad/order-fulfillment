package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// ------------------------------------------------------
// Create Order
// ------------------------------------------------------
func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}


func (r *Repository) CreateOrderTx(
	ctx context.Context,
	userID int64,
	addressID int64,
	// warehouseID *int64,
	orderNumber string,
	subtotal float64,
	shipping float64,
	tax float64,
	discount float64,
	total float64,
) (int64, error) {

	query := `
	INSERT INTO orders (
		user_id,
		address_id,
		order_number,
		status,
		subtotal,
		shipping_cost,
		tax,
		discount,
		total_amount
	)
	VALUES (
		$1,$2,$3,
		'pending',
		$4,$5,$6,$7,$8
	)
	RETURNING id
	`

	var orderID int64

	err := r.db.QueryRowContext(
		ctx,
		query,
		userID,
		addressID,
		orderNumber,
		subtotal,
		shipping,
		tax,
		discount,
		total,
	).Scan(&orderID)

	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	return orderID, nil
}

//
// ------------------------------------------------------
// Get Order By ID
// ------------------------------------------------------
//

func (r *Repository) GetOrderByID(
	ctx context.Context,
	orderID int64,
) (Order, error) {

	var order Order

	query := `
	SELECT
		id,
		user_id,
		address_id,
		order_number,
		status,
		subtotal,
		shipping_cost,
		tax,
		discount,
		total_amount,
		created_at,
		updated_at
	FROM orders
	WHERE id=$1
	`

	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&order.ID,
		&order.UserID,
		&order.AddressID,
		&order.OrderNumber,
		&order.Status,
		&order.Subtotal,
		&order.ShippingCost,
		&order.Tax,
		&order.Discount,
		&order.TotalAmount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return Order{}, ErrOrderNotFound
	}

	if err != nil {
		return Order{}, fmt.Errorf("failed to fetch order: %w", err)
	}

	return order, nil
}

//
// ------------------------------------------------------
// Get User Orders
// ------------------------------------------------------
//

func (r *Repository) GetOrdersByUser(
	ctx context.Context,
	userID int64,
) ([]Order, error) {

	query := `
	SELECT
		id,
		user_id,
		address_id,
		order_number,
		status,
		subtotal,
		shipping_cost,
		tax,
		discount,
		total_amount,
		created_at,
		updated_at
	FROM orders
	WHERE user_id=$1
	ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}
	defer rows.Close()

	orders := make([]Order, 0)

	for rows.Next() {

		var order Order

		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.AddressID,

			&order.OrderNumber,
			&order.Status,
			&order.Subtotal,
			&order.ShippingCost,
			&order.Tax,
			&order.Discount,
			&order.TotalAmount,
			&order.CreatedAt,
			&order.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

//
// ------------------------------------------------------
// Update Order Status
// ------------------------------------------------------
//

func (r *Repository) UpdateStatus(
	ctx context.Context,
	orderID int64,
	status string,
) error {

	query := `
	UPDATE orders
	SET
		status=$1,
		updated_at=NOW()
	WHERE id=$2
	`

	result, err := r.db.ExecContext(ctx, query, status, orderID)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrOrderNotFound
	}

	return nil
}


func (r *Repository)UpdateStatusTx(
    ctx context.Context,
    tx *sql.Tx,
    orderID int64,
    status string,
) error {

	query := `
	UPDATE orders
	SET
		status=$1,
		updated_at=NOW()
	WHERE id=$2
	`

	result, err := tx.ExecContext(ctx, query, status, orderID)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrOrderNotFound
	}

	return nil
}
//
// ------------------------------------------------------
// Delete Order
// ------------------------------------------------------
//

func (r *Repository) DeleteOrder(
	ctx context.Context,
	orderID int64,
) error {

	query := `
	DELETE FROM orders
	WHERE id=$1
	`

	result, err := r.db.ExecContext(ctx, query, orderID)
	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrOrderNotFound
	}

	return nil
}


func (r *Repository) CreateOrderItemTx(
	ctx context.Context,
	orderID int64,
	productID int64,
	warehouseID int64,
	quantity int,
	unitPrice float64,
	totalPrice float64,
) error {

	query := `
	INSERT INTO order_items
	(
		order_id,
		product_id,
		warehouse_id,
		quantity,
		unit_price,
		total_price
	)
	VALUES
	(
		$1,$2,$3,$4,$5,$6
	)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		orderID,
		productID,
		warehouseID,
		quantity,
		unitPrice,
		totalPrice,
	)

	if err != nil {
		return fmt.Errorf("failed to create order item: %w", err)
	}

	return nil
}

func (r *Repository) GetOrderItems(
	ctx context.Context,
	orderID int64,
) ([]OrderItem, error) {

	query := `
	SELECT
		id,
		order_id,
		product_id,
		warehouse_id,
		quantity,
		unit_price,
		total_price
	FROM order_items
	WHERE order_id = $1
	ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order items: %w", err)
	}
	defer rows.Close()

	items := make([]OrderItem, 0)

	for rows.Next() {

		var item OrderItem

		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.WarehouseID,
			&item.Quantity,
			&item.UnitPrice,
			&item.TotalPrice,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan order item: %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) DeleteOrderItemsTx(
	ctx context.Context,
	tx *sql.Tx,
	orderID int64,
) error {

	_, err := tx.ExecContext(
		ctx,
		`DELETE FROM order_items WHERE order_id = $1`,
		orderID,
	)

	if err != nil {
		return fmt.Errorf("failed to delete order items: %w", err)
	}

	return nil
}