package grpcclients

import (
	"inventory/internal/config"
	pb "inventory/proto/warehousepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectWarehouseGrpc(cfg config.Config) (pb.WarehouseProtoClient,error) {
	conn, err := grpc.NewClient(
		cfg.GrpcClient.ProductClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		return  nil,err
	}

	return  pb.NewWarehouseProtoClient(conn),nil
}