package main

import (
	"order/database"
	grpcclients "order/grpcClient"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"order"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "order/proto/orderpb"
)

// main server and then grpc server
func main(){

	cfg:= order.MustLoad()

	db:=database.ConnectDb(cfg)
    defer db.Close()

	userGrpcClient,err:= grpcclients.NewUserServiceClient(cfg)
	productGrpcClient,err:= grpcclients.NewProductGrpcClient(cfg)




	repo:= order.NewRepository(db)
	service:= order.NewService(
		repo,
		productGrpcClient,
		userGrpcClient,
	
	)
	handler:=order.NewHandler(service)
	grpcHandler :=order.NewGrpcHandler(service)
    // log.Print(grpcHandler)
	app:=gin.Default()

	cartRoute:=app.Group("/cart")
	order.RegisterRoutes(cartRoute,handler)
	


	grpcServ:= grpc.NewServer()
	
	pb.RegisterOrderServiceServer(
		grpcServ,
		grpcHandler,
	)
	

	lis,err := net.Listen("tcp",cfg.GrpcServ.Port)

	if err !=nil {
		log.Fatal("Error to net listen fail ",err)
	}

	go func (){
		log.Println("Product Grpc server starat on the port",cfg.Port)
		err:= grpcServ.Serve(lis)
		if err!=nil {
			log.Fatal("Error to start the grpcSercver",err)
		}
	}()


	serv:=http.Server{
		Addr: cfg.Addr,
		Handler: app,
	}

	done:=make(chan os.Signal,1)//make a channel

	signal.Notify(done,os.Interrupt,syscall.SIGTERM)
	
	go func ()  {
		log.Println("Server Start on the Addr",cfg.Addr)
		err:=serv.ListenAndServe()

	if err!= nil {
		log.Print("Error to Start the server",err)
	}
	}()

	<-done

	log.Println("server is sutting down !!!!")

	ctx ,cancle:= context.WithTimeout(context.Background(),time.Minute*5)

	defer cancle()

	err=serv.Shutdown(ctx)

	if err!= nil {
		log.Println("error to shutdownt the server",err)
		return
	}

	log.Println("server shut down successfullyyyyyy")

}