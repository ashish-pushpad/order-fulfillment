package grpc

import (
	"checkout/internal/config"
	pb "proto/inventory"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectInventory(cfg config.Config) pb.InventoryServiceClient {
	conn, err := grpc.NewClient(
		cfg.InventoryClient, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil { 
		log.Fatal(err)
	}

	return pb.NewInventoryServiceClient(conn)
}