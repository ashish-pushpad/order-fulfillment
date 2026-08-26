package main

import (
	"inventory/database/postgre"
	"inventory/internal"
	"inventory/internal/config"
	"inventory/internal/grpc"
	"net"

	"inventory/grpcClients"

	grpcServer "google.golang.org/grpc"

	pb "proto/inventory"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	cfg := config.MustLoad()
	godotenv.Load()
	db, err := postgre.ConnectDb(cfg)

	if err != nil {
		log.Fatalf("Fail to connect the db %s", err)
	}

	defer db.Close()

	productClient,err:=grpcclients.ConnectProductGrpc(cfg)
	if err != nil {
		log.Println("Error to conenct the ProductClient ",err)	
	}
	warehouClient,err:=grpcclients.ConnectWarehouseGrpc(cfg)
	if err != nil {
		log.Println("Error to conenct the warehosueClient ",err)	
	}
	

	r:=gin.Default()
	 
	api:=r.Group("inventory/v1")

	repo := internal.NewRepository(db)

	service := internal.NewService(repo,
		productClient,
		warehouClient,
	)

	httpHandler := internal.NewHandler(service)
	
	// inventoriesGroup := api.Group("/inventories")
	internal.RegisterRoutes(api,httpHandler)

	// Grpc Server Start here
	grpcSrv:=grpcServer.NewServer()// start the server 


	// get handler
	grpcHandler := grpc.NewServer(service)// all service provide to the handler so can user by handler 


	// regiter the handler to the server
	pb.RegisterInventoryServiceServer(// handler provide to the server (like handler ) where it can call the methods
		grpcSrv,
		grpcHandler,
	)


	lis, err := net.Listen("tcp", ":50051")// make a port to listen on the 50051
	if err != nil {
		log.Fatal(err)
	}

	go func ()  {
		slog.Info("gRPC Server started", "address", ":50051")// listent on the port 500051

	if err := grpcSrv.Serve(lis); err != nil {// server and list on the port 50051 
		log.Fatal(err)
	}
	}()


	// r := router.NewRouter(db)
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: r,// give the info about the router (handler have all the route and and service where to ridirect)
	}

	slog.Info("Server started", "address", cfg.Addr)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Fail to strart the server ", err)
		return
	}
}
