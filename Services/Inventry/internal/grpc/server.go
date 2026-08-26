package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"inventory/internal"
	pb "proto/inventory"
)

type Server struct {
	pb.UnimplementedInventoryServiceServer

	service *internal.Service
}

func NewServer(service *internal.Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) GetInventory(
	ctx context.Context,
	req *pb.GetInventoryRequest,
) (*pb.GetInventoryResponse, error) {

	stocks, err := s.service.GetAvailableStock(req.ProductId, req.WarehouseId)

	if err != nil {
		return nil, err
	}

	return &pb.GetInventoryResponse{
		Id:          1,
		ProductId:   req.ProductId,
		WarehouseId: req.WarehouseId,
		Quantity:    int32(stocks),
	}, nil

}

func (s *Server) ReserveStock(ctx context.Context, req *pb.ReserveStockRequest) (*emptypb.Empty, error) {
	err := s.service.ReserveStock(ctx, req.ProductId, req.WarehouseId, int(req.Quantity))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) FindWarehouseForProduct(ctx context.Context, req *pb.FindWarehouseForProductRequest) (*pb.FindWarehouseForProductResponse, error) {
	wareId, err := s.service.FindWarehouseForProduct(ctx, req.ProductId ,int(req.Quantity))

	if err != nil {
		return nil, err
	}
	return &pb.FindWarehouseForProductResponse{
		WarehouseId:wareId,
	}, nil
}
