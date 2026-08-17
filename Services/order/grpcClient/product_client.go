package grpcclient

import (
	"order"
	pb "order/proto/productpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductGrpcClient(cfg order.Config) (pb.ProductServiceClient,error) {
	conn, err := grpc.NewClient(
		cfg.GrpcClient.ProductClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		return  nil,err
	}

	return  pb.NewProductServiceClient(conn),nil
}