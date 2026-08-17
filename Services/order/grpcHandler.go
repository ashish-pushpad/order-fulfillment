package order

import (
	"context"
	"database/sql"
	"log"
	pb "order/proto/orderpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service *Service
}

func NewGrpcHandler(service *Service) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}




func (h *GrpcHandler) CreateOrder(ctx context.Context, req *pb.CreatOrderRequest) (*pb.CreatOrderResponse, error) {
	orderId, err := h.service.CreateOrder(
		ctx , 
		req.UserId , 
		req.AddressId ,
		req.OrderNumber,
		float64( req.Subtotal ),
		  float64(req.Shipping) ,
		   float64(req.Tax) ,
		    float64(req.Discount ),
			 float64(req.Total) ,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		log.Println("error to CReateOrdre",err)
		return nil, status.Error(codes.Internal, "internal server error")
	}


	return &pb.CreatOrderResponse{
		OrderId: orderId,
	}, nil
}




func (h *GrpcHandler) CreateOrderItem(ctx context.Context,req *pb.CreateOrderItemRequest)(*emptypb.Empty,error) {

	err:= h.service.CreateOrderItem(ctx,
		req.OrderId , req.ProductId ,
		req.WarehouseId , int(req.Quantity),
		 float64(req.UnitPrice) , float64(req.TotalPrice) ,
		 ) 
	
	if err!= nil {
		return  &emptypb.Empty{},err
	}

	return &emptypb.Empty{},nil

}
