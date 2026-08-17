package grpcclient

import (
	"log"
	pb "order/proto/userpb"
	"order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient(cfg order.Config)(pb.UserServiceClient ,error) {
	conn,err:=grpc.NewClient(
		cfg.UserClient, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err!=nil {
		log.Println("Error to connec to the user service",err)
		return nil,err
	}
	return  pb.NewUserServiceClient(conn),nil
}