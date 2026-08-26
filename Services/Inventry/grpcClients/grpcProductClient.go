package grpcclients

import (
	"inventory/internal/config"
	
	pb "proto/product"
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