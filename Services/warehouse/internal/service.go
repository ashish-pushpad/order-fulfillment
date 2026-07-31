package internal

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateWarehouse(warehouse CreateWarehouseBody) error {

	// Business rules

	// Example:
	// Validate name
	// Check duplicate email
	// Hash password
	// Send event

	return s.repo.CreateWarehouse(warehouse)
}

func (s *Service) GetWarehouse(ctx context.Context, id int64) (GetWarehouseResponse, error) {

	p, err := s.repo.GetWarehouse(ctx,id)

	if err != nil {
		return GetWarehouseResponse{}, err
	}

	warehouse := GetWarehouseResponse{
		Id:      p.Id,
		Name:    p.Name,
		City:    p.City,
		Address: p.Address,
	}

	return warehouse, nil
}

func (s *Service) GetWarehouses(limit, page *int) ([]GetWarehouseResponse, error) {

	warehouses, err := s.repo.GetWarehouses(limit, page)

	if err != nil {
		return []GetWarehouseResponse{}, err
	}

	return warehouses, nil
}

func (s *Service) UpdateWarehouse(id int64, data UpdateWarehouseBody) (UpdateWarehouseBodyRes, error) {

	warehouse, err := s.repo.UpdateWarehouse(id, data)

	if err != nil {
		return UpdateWarehouseBodyRes{}, err
	}

	return warehouse, nil

}

func (s *Service) DeleteWarehouse(id int64) (int64, error) {
	deletedWarehouse, err := s.repo.DeleteWarehouse(id)
	if err != nil {
		return 0, err
	}

	return deletedWarehouse, nil
}