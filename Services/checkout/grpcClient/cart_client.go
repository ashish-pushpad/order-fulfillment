package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	 pb "checkout/proto/cartpb"
)

func NewCartServiceClient() ( pb.CartServiceClient,error) {
		conn,err:=grpc.NewClient(
			"localhost:50056", 
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err!=nil {
			return nil ,err
		}

		return pb.NewCartServiceClient(conn),nil
}