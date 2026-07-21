package router

import (
	"database/sql"
	"inventory/internal"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {

	r:=gin.Default()
	 
	api:=r.Group("inventory/v1")

	inventoriesRepo:= internal.NewRepository(db)
	inventoriesServie:= internal.NewService(inventoriesRepo)
	inventoriesHandler:= internal.NewHandler(inventoriesServie)

	inventoriesGroup := api.Group("/inventories")

	

	internal.RegisterRoutes(inventoriesGroup,inventoriesHandler)



	return  r
}