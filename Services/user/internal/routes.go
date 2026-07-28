package internal

import (

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {

    rg.POST("/", h.CreateUser)
    rg.GET("/:id",h.GetUser)
    rg.PUT("/:id",h.UpdateUser)
    rg.DELETE("/:id",h.DeleteUser)


    // rg.POST("/", CreateUser)
}