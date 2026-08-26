package internal

import (
	"context"
	"database/sql"
	"fmt"
	product "proto/product"
	warehouse "proto/warehouse"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type Service struct {
	repo *Repository
	product product.ProductServiceClient 
	warehouse warehouse.WarehouseProtoClient
}

func NewService(repo *Repository,product product.ProductServiceClient,warehouse warehouse.WarehouseProtoClient ) *Service {
	return &Service{
		repo: repo,
		product: product,
		warehouse: warehouse,
	}
}
 

func (s *Service) CreateInventory(ctx context.Context,body CreateInventoryBody) error {

	// Future business logic:
	// - Check if product exists
	// - Check if warehouse exists
	// - Prevent duplicate inventory
	_, err := s.product.IsExistProduct(ctx, &product.IsExistProductRequest{
    Id: body.ProductID,
})

if err != nil {
    if status.Code(err) == codes.Unavailable {
        return fmt.Errorf("product service unavailable: %w", err)
    }

    return fmt.Errorf("product validation failed: %w", err)
}

_, err = s.warehouse.GetWarehouse(ctx, &warehouse.GetWarehouseRequest{
    Id: body.WarehouseID,
})

if err != nil {
    if status.Code(err) == codes.Unavailable {
        return fmt.Errorf("warehouse service unavailable: %w", err)
    }

    return fmt.Errorf("warehouse validation failed: %w", err)
}

	return s.repo.CreateInventory(body)
}

func (s *Service) GetInventory(id int64) (InventoryResponse, error) {

	inventory, err := s.repo.GetInventory(id)
	if err != nil {
		return InventoryResponse{}, err
	}

	return inventory, nil
}


func (s *Service) GetInventories(page, limit *int) ([]InventoryResponse, error) {

	inventories, err := s.repo.GetInventories(page, limit)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (s *Service) UpdateInventory(
	id int64,
	body UpdateInventoryBody,
) (InventoryResponse, error) {

	inventory, err := s.repo.UpdateInventory(id, body)
	if err != nil {
		return InventoryResponse{}, err
	}

	return inventory, nil
}

func (s *Service) DeleteInventory(id int64) error {

	return s.repo.DeleteInventory(id)
}



func (s *Service) ReserveStock(ctx context.Context, productID, warehouseID int64, qty int) error {
	if qty <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	return s.repo.ReserveStock(ctx  ,productID, warehouseID, qty)
}

func (s *Service) ReleaseStock(ctx context.Context,productID, warehouseID int64, qty int) error {

	if qty <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	return s.repo.ReleaseStockTx(ctx,productID, warehouseID, qty)
}

func (s *Service) CommitReservationTx(ctx context.Context, tx *sql.Tx,productID, warehouseID int64, qty int) error {

	if qty <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	return s.repo.CommitReservationTx(ctx,tx,productID, warehouseID, qty)
}

func (s *Service) GetAvailableStock(productID, warehouseID int64) (int, error) {

	return s.repo.GetAvailableStock(productID, warehouseID)
}

func (s *Service) FindWarehouseForProduct(
	ctx context.Context,
	productID int64,
	qty int,
) (int64, error) {

	if qty <= 0 {
		return 0, fmt.Errorf("quantity must be greater than zero")
	}

	return s.repo.FindWarehouseForProduct(
		ctx,
		productID,
		qty,
	)
}