package inventories

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {

    rg.POST("/", h.CreateInventory)
    rg.DELETE("/:id",h.DeleteInventory)
    rg.PUT("/:id",h.UpdateInventory)
    rg.GET("/stock/:warehouseId/products/:productId",h.GetAvailableStock)
    rg.GET("/",h.GetInventories)
    rg.GET("/:id",h.GetInventory)
}