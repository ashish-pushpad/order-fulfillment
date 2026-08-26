package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"
	pb "proto/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "product/proto/productpb"
	
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


func (h *GrpcHandler) GetProductsByIds(
    ctx context.Context,
    req *pb.GetProductsByIdsRequest,
) (*pb.GetProductsByIdsResponse, error) {

    products, err := h.service.repo.GetProductsByIds(ctx,req.Ids)
    if err != nil {
        switch {
        case errors.Is(err, ErrProductNotFound):
            return nil, status.Error(codes.NotFound, "product not found")

        default:
			log.Println("error at GetProductsByIds",err)
            return nil, status.Error(codes.Internal, "internal server error")
        }
    }

    pbProducts := make([]*pb.Product, 0, len(products))

	for _, p := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			ImgUrl:      p.ImgUrl,
		})
	}

	return &pb.GetProductsByIdsResponse{
		Products: pbProducts,
	}, nil
}


func (h *GrpcHandler) GetProductPriceById(ctx context.Context,req *pb.GetProductPriceByIdRequest) (*pb.GetProductPriceByIdResponse,error){
	product,err:=h.service.GetProductPriceById(ctx,req.Id)
	if err !=nil{
		log.Println("Error to get the price by id in product",err)
		return  &pb.GetProductPriceByIdResponse{},err
	}

	return  &pb.GetProductPriceByIdResponse{
		Id: product.Id,
		Price: float32(product.Price),
	},err
}