package grpcclient

import (
	"order"
	pb "order/proto/addresspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAddressServiceClient(cfg order.Config) ( pb.AddressServiceClient,error) {
		conn,err:=grpc.NewClient(
			cfg.GrpcClient.AddressClient, 
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err!=nil {
			return nil ,err
		}

		return pb.NewAddressServiceClient(conn),nil
}