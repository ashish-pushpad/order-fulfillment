package grpc

import (
	"log"
	"google.golang.org/grpc"
	 pb "checkout/proto/inventorypb"
	 "google.golang.org/grpc/credentials/insecure"
)

func ConnectInventory() pb.InventoryServiceClient {
	conn, err := grpc.NewClient(
		"localhost:50051", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil { 
		log.Fatal(err)
	}

	return pb.NewInventoryServiceClient(conn)
}