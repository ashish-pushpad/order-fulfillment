package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)


var  ErrProductNotFound = errors.New("product not found")


type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}




func (r *Repository) CreateProduct(Product CreateProductBody) error {

	_, err := r.db.Exec(`
		INSERT INTO products(name,description,price,img_url)
		VALUES($1,$2,$3,$4)
	`, Product.Name, Product.Description, Product.Price, Product.ImgUrl)

	return err
}

func (r *Repository) GetProduct(id int64) (GetProductResponse, error) {

	row := r.db.QueryRow(`
					SELECT id,name,price,description,img_url
					FROM Products
					WHERE id = $1
				`, id)

	var product GetProductResponse
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.ImgUrl)
	if err != nil {
		return GetProductResponse{}, err
	}

	return product, nil
}

func (r *Repository) GetProducts(page, limit *int) ([]Product, error) {
	fmt.Println(page == nil)
	fmt.Println(limit == nil)
	var (
		rows *sql.Rows
		err  error
	)
	fmt.Print("Im Get products")
	if page != nil && limit != nil {
		offset := (*page - 1) * (*limit)

		rows, err = r.db.Query(`
			SELECT id, name, price, description, img_url, created_at
			FROM products
			ORDER BY id
			LIMIT $1 OFFSET $2
		`, *limit, offset)
	} else {
		rows, err = r.db.Query(`
			SELECT id, name, price, description, img_url, created_at
			FROM products
			ORDER BY id
		`)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.ImageURL,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Print("Im Get products jsut befroe the return")
	slog.Debug("products fetched", "products", products)
	return products, nil
}

func (r *Repository) UpdateProduct(id int64, product UpdateProductBody) (UpdateProductBodyRes, error) {
	var updatedProduct UpdateProductBodyRes

	err := r.db.QueryRow(`
		UPDATE products
		SET
			name = $1,
			price = $2,
			description = $3,
			img_url = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING id, name, price, description, img_url
	`,
		product.Name,
		product.Price,
		product.Description,
		product.ImgUrl,
		id,
	).Scan(
		&updatedProduct.Id,
		&updatedProduct.Name,
		&updatedProduct.Price,
		&updatedProduct.Description,
		&updatedProduct.ImgUrl,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return UpdateProductBodyRes{}, fmt.Errorf("product not found")
	}

	if err != nil {
		return UpdateProductBodyRes{}, err
	}

	return updatedProduct, nil
}

func (r *Repository) DeleteProduct(id int64) (int64, error) {
	result, err := r.db.Exec(`
		DELETE FROM products
		WHERE id = $1
	`, id)

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return 0, fmt.Errorf("product not found")
	}

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}


func (r *Repository) IsExistProduct(id int64) ( int64, error) {

	row := r.db.QueryRow(`
					SELECT id
					FROM Products
					WHERE id = $1
				`, id)

	var productId int64

	err:=row.Scan(&productId)

	
	if err != nil {
		return 0, err
	}

	return productId, nil
}

func (r *Repository) GetProductsByIds (ctx context.Context,ids []int64) ([]GetProductsByIdsResponse,error){
	rows,err:= r.db.QueryContext(ctx , `
	SELECT id,name,description,img_url From products 
	WHERE id = ANY($1)
	`,ids)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []GetProductsByIdsResponse

	for rows.Next() {
		var product GetProductsByIdsResponse

		err:= rows.Scan(&product.Id,&product.Name,&product.Description,&product.ImgUrl)

		if  err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	if err := rows.Err(); err != nil {
        return nil, err
    }
	return  products,nil
}

func (r *Repository) GetProductPriceById (ctx context.Context,id int64 ) (GetProductPriceByIdResponse ,error){
	var product GetProductPriceByIdResponse
	row:= r.db.QueryRowContext(ctx , `
	SELECT id,price From products 
	WHERE id = $1
	`,id)
 
		err:= row.Scan(&product.Id,&product.Price)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return GetProductPriceByIdResponse{}, ErrProductNotFound
			}
			return GetProductPriceByIdResponse{}, err
		}

	
	return  product,nil
}