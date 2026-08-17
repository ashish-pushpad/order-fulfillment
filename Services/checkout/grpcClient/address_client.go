package grpc

import (
	"checkout/internal/config"
	pb "checkout/proto/addresspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAddressServiceClient(cfg config.Config) ( pb.AddressServiceClient,error) {
		conn,err:=grpc.NewClient(
			cfg.AddressClient, 
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err!=nil {
			return nil ,err
		}

		return pb.NewAddressServiceClient(conn),nil
}