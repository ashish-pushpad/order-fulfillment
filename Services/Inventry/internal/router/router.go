package router

import (
	"database/sql"
	"inventory/grpcclients"
	"inventory/internal"
	"inventory/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB,cfg config.Config) *gin.Engine {

	r:=gin.Default()
		
	productClient,err:=grpcclients.ConnectProductGrpc(cfg)
	if err != nil {
		log.Println("Error to conenct the ProductClient ",err)	
	}
	warehouClient,err:=grpcclients.ConnectWarehouseGrpc(cfg)
	if err != nil {
		log.Println("Error to conenct the warehosueClient ",err)	
	}
	
	api:=r.Group("inventory/v1")

	inventoriesRepo:= internal.NewRepository(db)
	inventoriesServie:= internal.NewService(inventoriesRepo,
		productClient,
		warehouClient,
)
	inventoriesHandler:= internal.NewHandler(inventoriesServie)

	inventoriesGroup := api.Group("/inventories")

	

	internal.RegisterRoutes(inventoriesGroup,inventoriesHandler)



	return  r
}