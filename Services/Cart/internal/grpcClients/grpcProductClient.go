package grpcclients

import (
    "google.golang.org/grpc"
	 pb "cartservice/proto/productpb"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectProductGrpc() (pb.ProductServiceClient,error) {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		return  nil,err
	}

	return  pb.NewProductServiceClient(conn),nil
}