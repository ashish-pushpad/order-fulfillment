package grpc

import (
	"checkout/internal/config"
	pb "proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewOrderServiceClient(cfg config.Config) ( pb.OrderServiceClient,error) {
		conn,err:=grpc.NewClient(
			cfg.OrderClient, 
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err!=nil {
			return nil ,err
		}

		return pb.NewOrderServiceClient(conn),nil
}