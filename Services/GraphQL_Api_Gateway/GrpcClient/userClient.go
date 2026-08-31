package grpcclient

import (
	"apigateway/internal/config"
	user "proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClient(cfg *config.Config) (user.UserServiceClient,error) {
	conn,err:=grpc.NewClient(
		cfg.UserClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err!=nil {
		return nil,err
	}

	return user.NewUserServiceClient(conn),nil

	
}