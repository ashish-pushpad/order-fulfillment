package grpcclients

import (
	"cartservice/internal/config"
	pb "cartservice/proto/productpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectProductGrpc(cfg config.Config) (pb.ProductServiceClient,error) {
	conn, err := grpc.NewClient(
		cfg.GrpcClient.ProductClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		return  nil,err
	}

	return  pb.NewProductServiceClient(conn),nil
}