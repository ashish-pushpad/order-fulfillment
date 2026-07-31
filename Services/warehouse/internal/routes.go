package internal

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {

    rg.POST("/", h.CreateWarehouse)
    rg.GET("/",h.GetWarehouses)
    rg.GET("/:id",h.GetWarehouse)
    rg.PUT("/:id",h.UpdateWarehouse)
    rg.DELETE("/:id",h.DeleteWarehouse)
    // rg.POST("/", CreateUser)
}