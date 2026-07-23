package internal

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {

    rg.POST("/", h.CreateProduct)
    rg.GET("/",h.GetProducts)
    rg.GET("/:id",h.GetProduct)
    rg.PUT("/:id",h.UpdateProduct)
    rg.DELETE("/:id",h.DeleteProduct)
}