package grpc

import (
	"context"
	"inventory/internal"
	pb "inventory/proto/inventorypb"
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

	stocks,err:=s.service.GetAvailableStock(req.ProductId,req.WarehouseId)

	if err!=nil{
		return  nil, err
	}

	return &pb.GetInventoryResponse{
		Id:          1,
		ProductId:   req.ProductId,
		WarehouseId: req.WarehouseId,
		Quantity:    int32(stocks),
	}, nil
    
}