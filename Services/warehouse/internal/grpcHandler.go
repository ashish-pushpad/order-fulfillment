package internal

import (
	pb "warehouse/proto/warehousepb"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	pb.UnimplementedWarehouseProtoServer
	service *Service
}

func NewGrpcHandler(service *Service) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (h *GrpcHandler) GetUserCart(ctx context.Context, req *pb.GetWarehouseRequest) (*pb.GetWarehouseResponse, error) {
	resp, err := h.service.GetWarehouse(ctx , req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "warehouse not found")
		}
		log.Println("error to getWArehouse",err)
		return nil, status.Error(codes.Internal, "internal server error")
	}


	return &pb.GetWarehouseResponse{
		Id:resp.Id,
		Name: resp.Name,
		City: resp.City,
		Addressd: resp.Address,
	}, nil
}


