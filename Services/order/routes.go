package order

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
  

    rg.GET("/:id", h.GetOrderByID)
    rg.GET("/user/:id",h.GetOrdersByUser)
    rg.GET("/item/:id",h.GetOrderItems)
    // cart.POST("/items", h.AddItem)
    // cart.PUT("/items/:id", h.UpdateQuantity)
    // cart.DELETE("/items/:id", h.RemoveItem)
    // cart.DELETE("/", h.EmptyCart)
}