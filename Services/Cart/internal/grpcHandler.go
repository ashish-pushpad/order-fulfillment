package internal

import (
	pb "cartservice/proto/cartpb"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcHandler struct {
	pb.UnimplementedCartServiceServer
	service *Service
}

func NewGrpcHandler(service *Service) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (h *GrpcHandler) GetUserCart(ctx context.Context, req *pb.GetUserCartRequest) (*pb.GetUserCartResponse, error) {
	cart, err := h.service.GetUserCart(ctx , req.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		log.Println("error to getUserCart",err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	items:= make([] *pb.CartItemResponse,0,len(cart.Items))

	for _,i:= range cart.Items {
		items = append(items, &pb.CartItemResponse{
			  Id : i.ID,
			  ProductId :i.ProductID,
			  ProductName :i.ProductName,
			  Quantity : int64(i.Quantity),
			  UnitPrice :float32(i.UnitPrice),
			  TotalPrice : float32(i.TotalPrice),
		} )
	}

	return &pb.GetUserCartResponse{
		Id: cart.ID,
		UserId:cart.UserID,
		Items:items,
		SubTotal: float32(cart.Subtotal) ,
		TotalItems:int64(cart.TotalItems),
	}, nil
}


func (h *GrpcHandler) EmptyCart(ctx context.Context,req *pb.EmptyCartRequest)(*emptypb.Empty,error) {

	err:= h.service.EmptyCart(ctx,req.UserId)	
	if err!= nil {
		return  &emptypb.Empty{},err
	}

	return &emptypb.Empty{},nil

}
