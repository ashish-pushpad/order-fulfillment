package internal

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

func (r *Repository) CreateWarehouse(Warehouse CreateWarehouseBody) error {

	_, err := r.db.Exec(`
		INSERT INTO warehouse(name,city,address)
		VALUES($1,$2,$3)
	`, Warehouse.Name, Warehouse.City, Warehouse.Address)
	return err
}

func (r *Repository) GetWarehouse(ctx context.Context,id int64) (GetWarehouseResponse, error) {


	row := r.db.QueryRowContext(ctx,`
					SELECT id,name,city,address
					FROM warehouse
					WHERE id = $1
				`, id)

	var warehouse GetWarehouseResponse
	err := row.Scan(&warehouse.Id,&warehouse.Name,&warehouse.City,&warehouse.Address)
	if err != nil {
		return GetWarehouseResponse{}, err
	}

	return warehouse, nil
}

func (r *Repository) GetWarehouses(page, limit *int) ([]GetWarehouseResponse, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if page != nil && limit != nil {
		offset := (*page - 1) * (*limit)

		rows, err = r.db.Query(`
			SELECT id, name, city, address
			FROM warehouse
			ORDER BY id
			LIMIT $1 OFFSET $2
		`, *limit, offset)
	} else {
		rows, err = r.db.Query(`
			SELECT id, name, city, address
			FROM warehouse
			ORDER BY id
		`)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []GetWarehouseResponse

	for rows.Next() {
		var warehouse GetWarehouseResponse

		err := rows.Scan(
			&warehouse.Id,
			&warehouse.Name,
			&warehouse.City,
			&warehouse.Address,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, warehouse)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) UpdateWarehouse(id int64, warehouse UpdateWarehouseBody) (UpdateWarehouseBodyRes, error) {
	var updatedWarehouse UpdateWarehouseBodyRes

	err := r.db.QueryRow(`
		UPDATE warehouse
		SET
			name = $1,
			city = $2,
			address = $3
		WHERE id = $4
		RETURNING id, name, city, address
	`,
		warehouse.Name,
		warehouse.City,
		warehouse.Address,
		id	).Scan(
		&updatedWarehouse.Id,
		&updatedWarehouse.Name,
		&updatedWarehouse.City,
		&updatedWarehouse.Address,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return UpdateWarehouseBodyRes{}, fmt.Errorf("warehouse not found")
	}

	if err != nil {
		return UpdateWarehouseBodyRes{}, err
	}

	return updatedWarehouse, nil
}


func (r *Repository) DeleteWarehouse(id int64) (int64,error) {
	result, err := r.db.Exec(`
		DELETE FROM warehouse
		WHERE id = $1
	`, id)


	rowsAffected, err := result.RowsAffected()


	if rowsAffected == 0 {
		return 0,fmt.Errorf("warehouse not found")
	}

	if err != nil {
		return 0,err
	}

	return rowsAffected,nil
}
