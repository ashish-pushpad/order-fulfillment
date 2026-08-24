package main

import (
	grpcclient "apigateway/GrpcClient"
	config "apigateway/internal/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"apigateway/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)


func main(){
	cfg:= config.MustLoad()

	r:=gin.Default()


    userClient, err := grpcclient.UserClient(cfg)
    if err != nil {
        log.Println("error to connect the userClient  ",err)
    }

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserClient: userClient,
	}}))

	

	
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.GET("/",gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	// http.Handle("/query", srv)
	r.POST("/query",gin.WrapH(srv))

	// r.POST("/graphql",func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK,gin.H{
	// 		"message":"App is running",
	// 	})	
	// })

	serv:=http.Server{
		Addr: cfg.Addr,
		Handler: r,
	}


	log.Println("sdrver Started on the port ",cfg.Addr)
	log.Fatal(serv.ListenAndServe())

	
}