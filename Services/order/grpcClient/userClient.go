package grpcclient

import (
	"log"
	pb "order/proto/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserServiceClient()(pb.UserServiceClient ,error) {
	conn,err:=grpc.NewClient(
		"localhost:50054", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err!=nil {
		log.Println("Error to connec to the user service",err)
		return nil,err
	}
	return  pb.NewUserServiceClient(conn),nil
}