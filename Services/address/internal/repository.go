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

func (r *Repository) CreateAddress(userID int64, address CreateAddressBody) error {
	fmt.Println("userId in add addrs",userID)


	_, err := r.db.Exec(`
		INSERT INTO address (
			full_name,
			phone_number,
			house_no,
			apartment,
			colony,
			city,
			state,
			user_id,
			pin_code,
			is_default
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`,
		address.FullName,
		address.PhoneNumber,
		address.HouseNo,
		address.Apartment,
		address.Colony,
		address.City,
		address.State,
		userID,
		address.PinCode,
		address.IsDefault,
	)

	return err
}

func (r *Repository) GetAddress(ctx context.Context,id int64) (AddressResponse, error) {
	var address AddressResponse

	err := r.db.QueryRowContext(ctx,`
		SELECT
			id,
			full_name,
			phone_number,
			house_no,
			apartment,
			colony,
			city,
			state,
			user_id,
			pin_code,
			is_default
		FROM address
		WHERE id = $1
	`, id).Scan(
		&address.Id,
		&address.FullName,
		&address.PhoneNumber,
		&address.HouseNo,
		&address.Apartment,
		&address.Colony,
		&address.City,
		&address.State,
		&address.UserId,
		&address.PinCode,
		&address.IsDefault,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return AddressResponse{}, fmt.Errorf("address not found")
	}

	if err != nil {
		return AddressResponse{}, err
	}

	return address, nil
}

func (r *Repository) GetAddresses(page, limit *int) ([]AddressResponse, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if page != nil && limit != nil {
		offset := (*page - 1) * (*limit)

		rows, err = r.db.Query(`
			SELECT
				id,
				full_name,
				phone_number,
				house_no,
				apartment,
				colony,
				city,
				state,
				user_id,
				pin_code,
				is_default
			FROM address
			ORDER BY id
			LIMIT $1 OFFSET $2
		`, *limit, offset)
	} else {
		rows, err = r.db.Query(`
			SELECT
				id,
				full_name,
				phone_number,
				house_no,
				apartment,
				colony,
				city,
				state,
				user_id,
				pin_code,
				is_default
			FROM address
			ORDER BY id
		`)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []AddressResponse

	for rows.Next() {
		var address AddressResponse

		err := rows.Scan(
			&address.Id,
			&address.FullName,
			&address.PhoneNumber,
			&address.HouseNo,
			&address.Apartment,
			&address.Colony,
			&address.City,
			&address.State,
			&address.UserId,
			&address.PinCode,
			&address.IsDefault,
		)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *Repository) UpdateAddress(id int64, address UpdateAddressBody) (AddressResponse, error) {
	var updatedAddress AddressResponse

	err := r.db.QueryRow(`
		UPDATE address
		SET
			full_name = $1,
			phone_number = $2,
			house_no = $3,
			apartment = $4,
			colony = $5,
			city = $6,
			state = $7,
			pin_code = $8,
			is_default = $9
		WHERE id = $10
		RETURNING
			id,
			full_name,
			phone_number,
			house_no,
			apartment,
			colony,
			city,
			state,
			user_id,
			pin_code,
			is_default
	`,
		address.FullName,
		address.PhoneNumber,
		address.HouseNo,
		address.Apartment,
		address.Colony,
		address.City,
		address.State,
		address.PinCode,
		address.IsDefault,
		id,
	).Scan(
		&updatedAddress.Id,
		&updatedAddress.FullName,
		&updatedAddress.PhoneNumber,
		&updatedAddress.HouseNo,
		&updatedAddress.Apartment,
		&updatedAddress.Colony,
		&updatedAddress.City,
		&updatedAddress.State,
		&updatedAddress.UserId,
		&updatedAddress.PinCode,
		&updatedAddress.IsDefault,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return AddressResponse{}, fmt.Errorf("address not found")
	}

	if err != nil {
		return AddressResponse{}, err
	}

	return updatedAddress, nil
}

func (r *Repository) DeleteAddress(id int64) error {
	result, err := r.db.Exec(`
		DELETE FROM address
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
		return fmt.Errorf("address not found")
	}

	return nil
}