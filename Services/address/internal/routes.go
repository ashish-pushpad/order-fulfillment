package internal

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {

    rg.POST("/", h.CreateAddress)
    rg.GET("/:id",h.GetAddress)
    rg.GET("/",h.GetAddresses)
    // rg.GET("/",h.CreateAddress)
    rg.PUT("/:id",h.UpdateAddress)
    rg.DELETE("/:id",h.DeleteAddress)
    // rg.POST("/", CreateUser)
}