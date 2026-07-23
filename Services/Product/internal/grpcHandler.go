package internal

import (
	"context"
	"database/sql"
	pb "product/proto/productpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	pb.UnimplementedProductServiceServer
	service *Service
}

func NewGrpcHandler(service *Service) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (h *GrpcHandler) IsExistProduct(ctx context.Context, req *pb.IsExistProductRequest) (*pb.IsExistProductResponse, error) {
	id, err := h.service.IsExistProduct(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &pb.IsExistProductResponse{
		Id: int64(id),
	}, nil
}
