package inventories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateInventory(body CreateInventoryBody) error {

	_, err := r.db.Exec(`
		INSERT INTO inventories
		(
			product_id,
			warehouse_id,
			quantity,
			reserved_quantity
		)
		VALUES ($1,$2,$3,0)
	`,
		body.ProductID,
		body.WarehouseID,
		body.Quantity,
	)

	return err
}

func (r *Repository) GetInventory(id int64) (InventoryResponse, error) {

	var inventory InventoryResponse

	err := r.db.QueryRow(`
		SELECT
			id,
			product_id,
			warehouse_id,
			quantity,
			reserved_quantity
		FROM inventories
		WHERE id = $1
	`, id).Scan(
		&inventory.ID,
		&inventory.ProductID,
		&inventory.WarehouseID,
		&inventory.Quantity,
		&inventory.ReservedQuantity,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return InventoryResponse{}, fmt.Errorf("inventory not found")
	}

	if err != nil {
		return InventoryResponse{}, err
	}

	inventory.AvailableQuantity =
		inventory.Quantity - inventory.ReservedQuantity

	return inventory, nil
}

// ----------------------------
// GET ALL
// ----------------------------

func (r *Repository) GetInventories(page, limit *int) ([]InventoryResponse, error) {

	var (
		rows *sql.Rows
		err  error
	)

	if page != nil && limit != nil {

		offset := (*page - 1) * (*limit)

		rows, err = r.db.Query(`
			SELECT
				id,
				product_id,
				warehouse_id,
				quantity,
				reserved_quantity
			FROM inventories
			ORDER BY id
			LIMIT $1 OFFSET $2
		`, *limit, offset)

	} else {

		rows, err = r.db.Query(`
			SELECT
				id,
				product_id,
				warehouse_id,
				quantity,
				reserved_quantity
			FROM inventories
			ORDER BY id
		`)

	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var inventories []InventoryResponse

	for rows.Next() {

		var inventory InventoryResponse

		err := rows.Scan(
			&inventory.ID,
			&inventory.ProductID,
			&inventory.WarehouseID,
			&inventory.Quantity,
			&inventory.ReservedQuantity,
		)

		if err != nil {
			return nil, err
		}

		inventory.AvailableQuantity =
			inventory.Quantity - inventory.ReservedQuantity

		inventories = append(inventories, inventory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inventories, nil
}

// ----------------------------
// UPDATE
// ----------------------------

func (r *Repository) UpdateInventory(
	id int64,
	body UpdateInventoryBody,
) (InventoryResponse, error) {

	var inventory InventoryResponse

	err := r.db.QueryRow(`
		UPDATE inventories
		SET
			quantity = $1
		WHERE id = $2

		RETURNING
			id,
			product_id,
			warehouse_id,
			quantity,
			reserved_quantity
	`,
		body.Quantity,
		id,
	).Scan(
		&inventory.ID,
		&inventory.ProductID,
		&inventory.WarehouseID,
		&inventory.Quantity,
		&inventory.ReservedQuantity,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return InventoryResponse{}, fmt.Errorf("inventory not found")
	}

	if err != nil {
		return InventoryResponse{}, err
	}

	inventory.AvailableQuantity =
		inventory.Quantity - inventory.ReservedQuantity

	return inventory, nil
}

// ----------------------------
// DELETE
// ----------------------------

func (r *Repository) DeleteInventory(id int64) error {

	result, err := r.db.Exec(`
		DELETE FROM inventories
		WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("inventory not found")
	}

	return nil
}

// add the ttl so can expire it manually if there error happen
func (r *Repository) ReserveStockTx(
	ctx context.Context,
	tx *sql.Tx,
	productID int64,
	warehouseID int64,
	qty int,
) error {

	result, err :=tx.ExecContext(
		ctx,
		`
		UPDATE inventories
		SET reserved_quantity = reserved_quantity + $1
		WHERE product_id = $2
		AND warehouse_id = $3
		AND (quantity - reserved_quantity) >= $1
	`, qty, productID, warehouseID)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("not enough inventory")
	}

	return nil
}

func (r *Repository) ReleaseStockTx(
	ctx context.Context,
	tx *sql.Tx,
	productID int64,
	warehouseID int64,
	qty int,
) error {

	result, err := tx.ExecContext(
		ctx,
		`
    UPDATE inventories
    SET reserved_quantity = reserved_quantity - $1
    WHERE product_id = $2
      AND warehouse_id = $3
      AND reserved_quantity >= $1
    `,
		qty,
		productID,
		warehouseID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("invalid reservation")
	}

	return nil
}

func (r *Repository) CommitReservationTx(
	ctx context.Context,
	tx *sql.Tx,
	productID int64,
	warehouseID int64,
	qty int,
) error {

	result, err := tx.ExecContext(
		ctx,
		`
		UPDATE inventories
		SET
			quantity = quantity - $1,
			reserved_quantity = reserved_quantity - $1
		WHERE product_id = $2
		AND warehouse_id = $3
		AND reserved_quantity >= $1
	`, qty, productID, warehouseID)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("reservation not found")
	}

	return nil
}

func (r *Repository) GetAvailableStock(productID, warehouseID int64) (int, error) {

	var available int

	err := r.db.QueryRow(`
		SELECT quantity - reserved_quantity
		FROM inventories
		WHERE product_id = $1
		AND warehouse_id = $2
	`, productID, warehouseID).Scan(&available)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("inventory not found")
	}

	if err != nil {
		return 0, err
	}

	return available, nil
}


func (r *Repository) FindWarehouseForProduct(
	ctx context.Context,
	productID int64,
	qty int,
) (int64, error) {

	var warehouseID int64

	err := r.db.QueryRowContext(
		ctx,
		`
		SELECT warehouse_id
		FROM inventories
		WHERE product_id = $1
		  AND (quantity - reserved_quantity) >= $2
		ORDER BY (quantity - reserved_quantity) DESC
		LIMIT 1
		`,
		productID,
		qty,
	).Scan(&warehouseID)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("no warehouse has enough stock")
	}

	if err != nil {
		return 0, fmt.Errorf("failed to find warehouse: %w", err)
	}

	return warehouseID, nil
}