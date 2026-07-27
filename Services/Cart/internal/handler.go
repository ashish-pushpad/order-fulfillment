package internal


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


type AddItemRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int   `json:"quantity" binding:"required"`
}

type UpdateQuantityRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}



func (h *Handler) GetCart(c *gin.Context) {
	userID, err := h.getUserIDFromRequest(c)
	log.Println("userID in the Cart",userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	cartResponse, err := h.service.GetUserCart(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, ErrCartExpired) {
			c.JSON(http.StatusGone, gin.H{
				"error": "your cart has expired and has been cleared",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cartResponse)
}

func (h *Handler) AddItem(c *gin.Context) {
	// userID, err := h.getUserIDFromRequest(c)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "unauthorized",
	// 	})
	// 	return
	// }
	var userID int64
	userID=5
	var payload AddItemRequest
	log.Println(payload)
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.service.AddItemToCart(
		c.Request.Context(),
		userID,
		payload.ProductID,
		payload.Quantity,
	)

	if err != nil {
		if errors.Is(err, ErrInvalidQuantity) {
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
		"status": "item added successfully",
	})
}

func (h *Handler) UpdateQuantity(c *gin.Context) {
	userID, err := h.getUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	cartItemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid cart item id",
		})
		return
	}

	var payload UpdateQuantityRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = h.service.UpdateItemQuantity(
		c.Request.Context(),
		userID,
		cartItemID,
		payload.Quantity,
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidQuantity):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		case errors.Is(err, ErrItemNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": "c",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "quantity updated successfully",
	})
}


func (h *Handler) RemoveItem(c *gin.Context) {
	userID, err := h.getUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	cartItemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid cart item id",
		})
		return
	}

	err = h.service.RemoveItemFromCart(
		c.Request.Context(),
		userID,
		cartItemID,
	)

	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "cart item not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "item removed successfully",
	})
}


func (h *Handler) EmptyCart(c *gin.Context) {
	userID, err := h.getUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	err = h.service.EmptyCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "cart cleared successfully",
	})
}


func (h *Handler) getUserIDFromRequest(c *gin.Context) (int64, error) {
	value, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user id not found")
	}

	userID, ok := value.(int64)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	return userID, nil
}