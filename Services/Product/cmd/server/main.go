package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"product/database"
	"product/internal/config"
	"syscall"
	"time"
	"github.com/joho/godotenv"

	"product/internal"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	pb "product/proto/productpb"
)



// main server and then grpc server
func main(){
	err:=godotenv.Load()
	if err!=nil {
		log.Println("Error to load the env",err)
	}
	cfg:=config.MustLoad()

	db:=database.ConnectDb(cfg)
    defer db.Close()

	repo:= internal.NewRepository(db)
	service:=internal.NewService(repo)
	handler:=internal.NewHandler(service)
	grpcHandler :=internal.NewGrpcHandler(service)

	app:=gin.Default()

	prodcutRoute:=app.Group("/product")
	internal.RegisterRoutes(prodcutRoute,handler)
	


	grpcServ:= grpc.NewServer()
	
	pb.RegisterProductServiceServer(
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