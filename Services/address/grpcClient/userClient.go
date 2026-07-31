package grpcclient

import (
	"log"
	user "address/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserGrpcClient() (user.UserServiceClient,error) {
	conn,err:=grpc.NewClient(
		"localhost:50055",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		log.Println("error to connect the NewUserGrpcClient")
		return nil,err
	}

	return user.NewUserServiceClient(conn),nil

}
