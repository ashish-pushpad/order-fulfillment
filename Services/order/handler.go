package order

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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

// ------------------------------------------------------
// GET /orders/:id
// ------------------------------------------------------

func (h *Handler) GetOrderByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	order, err := h.service.GetOrderByID(c.Request.Context(), id)
	if err != nil {

		if errors.Is(err, ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// ------------------------------------------------------
// GET /users/:user_id/orders
// ------------------------------------------------------

func (h *Handler) GetOrdersByUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("error to pars the id in the getOr",err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	orders, err := h.service.GetOrdersByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// ------------------------------------------------------
// PATCH /orders/:id/status
// ------------------------------------------------------

func (h *Handler) UpdateStatus(c *gin.Context) {

	orderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	var body struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.UpdateStatus(
		c.Request.Context(),
		orderID,
		body.Status,
	)

	if err != nil {

		if errors.Is(err, ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, ErrInvalidStatus) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order status updated successfully",
	})
}

// ------------------------------------------------------
// DELETE /orders/:id
// ------------------------------------------------------

func (h *Handler) DeleteOrder(c *gin.Context) {

	orderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	err = h.service.DeleteOrder(c.Request.Context(), orderID)
	if err != nil {

		if errors.Is(err, ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order deleted successfully",
	})
}

// ------------------------------------------------------
// GET /orders/:id/items
// ------------------------------------------------------

func (h *Handler) GetOrderItems(c *gin.Context) {

	orderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	items, err := h.service.GetOrderItems(
		c.Request.Context(),
		orderID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, items)
}