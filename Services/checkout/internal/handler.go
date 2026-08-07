package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Checkout(c *gin.Context) {

	userID, ok := c.Get("user_id")
	if !ok {
		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"error": "unauthorized",
		// })
		// return
		userID=int64(1)
	}

	var body CheckoutBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.service.Checkout(
		c.Request.Context(),
		userID.(int64),
		body,
	)
	if err != nil {
		log.Println("Error in the checkout",err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}