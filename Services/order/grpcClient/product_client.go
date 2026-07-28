package grpcclient

import (
    "google.golang.org/grpc"
	 pb "order/proto/productpb"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductGrpcClient() (pb.ProductServiceClient,error) {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		return  nil,err
	}

	return  pb.NewProductServiceClient(conn),nil
}