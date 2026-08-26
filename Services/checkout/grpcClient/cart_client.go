package grpc

import (
	"checkout/internal/config"
	pb "proto/cart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCartServiceClient(cfg config.Config) ( pb.CartServiceClient,error) {
		conn,err:=grpc.NewClient(
			cfg.CartClient, 
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err!=nil {
			return nil ,err
		}

		return pb.NewCartServiceClient(conn),nil
}