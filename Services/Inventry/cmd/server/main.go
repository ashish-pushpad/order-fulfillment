package main

import (
	"inventory/database/postgre"
	"inventory/internal"
	"inventory/internal/config"
	"inventory/internal/grpc"
	  "net"

    grpcServer "google.golang.org/grpc"

    pb "inventory/proto/inventorypb"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.MustLoad()

	db, err := postgre.ConnectDb(cfg)

	if err != nil {
		log.Fatalf("Fail to connect the db %s", err)
	}

	defer db.Close()

	r:=gin.Default()
	 
	api:=r.Group("inventory/v1")

	repo := internal.NewRepository(db)

	service := internal.NewService(repo)

	httpHandler := internal.NewHandler(service)
	
	// inventoriesGroup := api.Group("/inventories")

	

	internal.RegisterRoutes(api,httpHandler)

	grpcHandler := grpc.NewServer(service)



		grpcSrv:=grpcServer.NewServer()
		pb.RegisterInventoryServiceServer(
    grpcSrv,
    grpcHandler,
)


lis, err := net.Listen("tcp", ":50051")
if err != nil {
    log.Fatal(err)
}

go func ()  {
	slog.Info("gRPC Server started", "address", ":50051")

if err := grpcSrv.Serve(lis); err != nil {
    log.Fatal(err)
}
}()


	// r := router.NewRouter(db)
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}

	slog.Info("Server started", "address", cfg.Addr)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Fail to strart the server ", err)
		return
	}
}
