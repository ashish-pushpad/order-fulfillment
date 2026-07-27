package internal

import (
	"context"
	"log"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateProduct(product CreateProductBody) error {

	// Business rules

	// Example:
	// Validate name
	// Check duplicate email
	// Hash password
	// Send event

	return s.repo.CreateProduct(product)
}

func (s *Service) GetProduct(id int64) (GetProductResponse, error) {

	p, err := s.repo.GetProduct(id)

	if err != nil {
		return GetProductResponse{}, err
	}

	product := GetProductResponse{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		ImgUrl:      p.ImgUrl,
	}

	return product, nil
}

func (s *Service) GetProducts(limit, page *int) ([]Product, error) {

	products, err := s.repo.GetProducts(limit, page)

	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *Service) UpdateProduct(id int64, data UpdateProductBody) (UpdateProductBodyRes, error) {

	product, err := s.repo.UpdateProduct(id, data)

	if err != nil {
		return UpdateProductBodyRes{}, err
	}

	return product, nil

}

func (s *Service) DeleteProduct(id int64) (int64, error) {
	deletedProduct, err := s.repo.DeleteProduct(id)
	if err != nil {
		return 0, err
	}

	return deletedProduct, nil
}

func (s *Service) IsExistProduct(ctx context.Context,id int64) (int64, error) {
	log.Println("Call the service isExistProduct")
	id, err := s.repo.IsExistProduct(id)
	
	if err != nil {
		return 0, err
	}
	return id, nil
}


// func (s *Service) GetProductPriceById (ctx context.Context , Ids []int64) ([]GetProductsByIdsResponse,error){
// 	products,err:= s.repo.GetProductsByIds(ctx,Ids)
	
// 	if err!=nil {
// 		return  nil,err
// 	}
// 	return  products,nil
// }

func (s *Service) GetProductPriceById (ctx context.Context, id int64)( GetProductPriceByIdResponse,error){
	product,err:= s.repo.GetProductPriceById(ctx,id)
	
		
	if err!=nil {
		return  GetProductPriceByIdResponse{},err
	}
	return  product,nil
}