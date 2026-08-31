package grpcclient

import (
	config "apigateway/internal/config"
	product "proto/product"


	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProductClient(cfg config.Config) (product.ProductServiceClient,error) {
	conn ,err:= grpc.NewClient(
		cfg.ProductClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err!= nil {
		return  nil,err
	}

	return  product.NewProductServiceClient(conn),nil
}