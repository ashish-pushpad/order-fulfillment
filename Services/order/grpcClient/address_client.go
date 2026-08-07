package grpcclient

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	 pb "order/proto/addresspb"
)

func NewAddressServiceClient() ( pb.AddressServiceClient,error) {
		conn,err:=grpc.NewClient(
			"localhost:50057", 
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err!=nil {
			return nil ,err
		}

		return pb.NewAddressServiceClient(conn),nil
}