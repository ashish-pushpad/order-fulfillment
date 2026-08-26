package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "proto/product"
	userpb "proto/user"
)

var (
	ErrInvalidAmount = errors.New("amount cannot be negative")
	ErrInvalidStatus = errors.New("invalid order status")
	ErrInvalidProduct = errors.New("Invalid Product Id")
)

type Service struct {
	repo *Repository
	product pb.ProductServiceClient
	user userpb.UserServiceClient
}

func NewService(repo *Repository,product pb.ProductServiceClient,user userpb.UserServiceClient ) *Service {
	return &Service{
		repo: repo,
		product:product,
		user:user,
	}
}

func (s *Service) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return s.repo.BeginTx(ctx)
}

func (s *Service) CreateOrder(
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

	if subtotal < 0 || shipping < 0 || tax < 0 || discount < 0 || total < 0 {
		return 0, ErrInvalidAmount
	}

	_,err:=s.user.GetUserById(ctx,&userpb.GetUserByIdRequest{Id: userID})
	if err!=nil {
		return  0,err
	}

	return s.repo.CreateOrderTx(
		ctx,
		userID,
		addressID,
		orderNumber,
		subtotal,
		shipping,
		tax,
		discount,
		total,
	)
}

func (s *Service) GetOrderByID(
	ctx context.Context,
	orderID int64,
) (Order, error) {

	return s.repo.GetOrderByID(ctx, orderID)
}

func (s *Service) GetOrdersByUser(
	ctx context.Context,
	userID int64,
) ([]Order, error) {

	return s.repo.GetOrdersByUser(ctx, userID)
}

func (s *Service) UpdateStatus(
	ctx context.Context,
	orderID int64,
	status string,
) error {

	switch status {
	case "pending", "confirmed", "cancelled", "shipped", "delivered":
	default:
		return ErrInvalidStatus
	}

	return s.repo.UpdateStatus(ctx, orderID, status)
}


func (s *Service) UpdateStatusTx(
	ctx context.Context,
	tx *sql.Tx,
	orderID int64,
	status string,
) error {

	switch status {
	case "pending", "confirmed", "cancelled", "shipped", "delivered":
	default:
		return ErrInvalidStatus
	}

	return s.repo.UpdateStatusTx(
		ctx,
		tx,
		orderID,
		status,
	)
}

func (s *Service) DeleteOrder(
	ctx context.Context,
	orderID int64,
) error {

	return s.repo.DeleteOrder(ctx, orderID)
}

func (s *Service) CreateOrderItem(
	ctx context.Context,
	orderID int64,
	productID int64,
	warehouseID int64,
	quantity int,
	unitPrice float64,
	totalPrice float64,
) error {

	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	if unitPrice < 0 || totalPrice < 0 {
		return ErrInvalidAmount
	}

	_,err:= s.product.IsExistProduct(ctx, &pb.IsExistProductRequest{
		Id: productID,
	})

	if err!=nil {
		return  err
	}

	return s.repo.CreateOrderItemTx(
		ctx,
		orderID,
		productID,
	    warehouseID,
		quantity,
		unitPrice,
		totalPrice,
	)
}

func (s *Service) GetOrderItems(
	ctx context.Context,
	orderID int64,
) ([]OrderItem, error) {

	return s.repo.GetOrderItems(ctx, orderID)
}

func (s *Service) DeleteOrderItems(
	ctx context.Context,
	tx *sql.Tx,
	orderID int64,
) error {

	return s.repo.DeleteOrderItemsTx(ctx, tx, orderID)
}

func (s *Service) Commit(tx *sql.Tx) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}
	return tx.Commit()
}

func (s *Service) Rollback(tx *sql.Tx) error {
	if tx == nil {
		return nil
	}
	return tx.Rollback()
}