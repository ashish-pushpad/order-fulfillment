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

func (r *Repository) CreateUser(user CreateUserBody) error {

	_, err := r.db.Exec(`
		INSERT INTO users(name,email,password_hash)
		VALUES($1,$2,$3)
	`, user.Name, user.Email, user.Password)

	return err
}

func (r *Repository) GetUser(ctx context.Context,id int64) (GetUserResponse, error) {


	row := r.db.QueryRowContext(
					ctx,`
					SELECT id,name, email
					FROM users
					WHERE id = $1
				`, id)

	var user GetUserResponse
	err := row.Scan(&user.Id,&user.Name,&user.Email)
	if err != nil {
		return GetUserResponse{}, err
	}

	return user, nil
}


func (r *Repository) UpdateUser(id int64, user UpdateUserBody) (User, error) {
	var updatedUser User

	err := r.db.QueryRow(`
		UPDATE users
		SET
			name = $1,
			email = $2,
			password = $3
		WHERE id = $4
		RETURNING id, name, email
	`, user.Name, user.Email, user.Password, id).Scan(
		&updatedUser.ID,
		&updatedUser.Name,
		&updatedUser.Email,
	)
    if errors.Is(err, sql.ErrNoRows) {
    return User{}, fmt.Errorf("user not found")
}
	if err != nil {
		return User{}, err
	}

	return updatedUser, nil
}

func (r *Repository) DeleteUser ( id int64) ( User,error){
		var deletedUser User
		err := r.db.QueryRow(`
		DELETE FROM users
		WHERE id = $1
		RETURNING id, name, email
	`, id).Scan(
		&deletedUser.ID,
		&deletedUser.Name,
		&deletedUser.Email,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, fmt.Errorf("user not found")
	}

	if err != nil {
		return User{}, err
	}

	return deletedUser, nil
}

func (r *Repository ) VerifyUser (email string ,password string) (GetLoginResponse, error){
	var user GetLoginResponse
	row:=r.db.QueryRow(`
	  SELECT id ,role FROM users WHERE email=$,hash_password=$
	  `,email,password)

	err := row.Scan(&user.Id,&user.Role)
	if err != nil {
		return GetLoginResponse{}, err
	}

	return user, nil
	}
