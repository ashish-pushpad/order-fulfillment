package internal

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
  

    rg.GET("/", h.GetCart)
    rg.POST("/items", h.AddItem)
    rg.PUT("/items/:id", h.UpdateQuantity)
    rg.DELETE("/items/:id", h.RemoveItem)
    rg.DELETE("/", h.EmptyCart)
}