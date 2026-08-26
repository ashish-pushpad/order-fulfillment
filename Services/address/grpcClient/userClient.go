package grpcclient

import (
	"address/internal/config"
	user "proto/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserGrpcClient(cfg config.Config) (user.UserServiceClient,error) {
	conn,err:=grpc.NewClient(
		cfg.GrpcClient.UserClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!=nil {
		log.Println("error to connect the NewUserGrpcClient")
		return nil,err
	}

	return user.NewUserServiceClient(conn),nil

}
