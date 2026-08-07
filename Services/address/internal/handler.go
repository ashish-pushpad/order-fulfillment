package internal

import (
	// "errors"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	// "theapp/internal/auth"

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

func (h *Handler) CreateAddress(c *gin.Context) {
	var req CreateAddressBody
    ctx:=c.Request.Context()
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind json", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	value,_:=c.Get("user_id")
	 log.Println("Usae id ",value)
	userId, ok := value.(int64)
	if !ok {
		// slog.Error("User id not provide", "error", err)
				
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": errors.New("jwt not valid"),
		// })
		// return
		userId=int64(1)
	}

	err := h.service.CreateAddress(ctx,userId, req)
	if err != nil {
		slog.Error("failed to create address", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "address created successfully",
		"address": req,
	})
}

func (h *Handler) GetAddress(c *gin.Context) {
	idParam := c.Param("id")
	ctx:= c.Request.Context()
	// userId, err := auth.GetUserID(c)
	var userId int64 =1;
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid address id",
		})
		return
	}

	address, err := h.service.GetAddress(ctx,userId, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, address)
}

func (h *Handler) GetAddresses(c *gin.Context) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	var (
		pagePtr  *int
		limitPtr *int
	)

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid page",
			})
			return
		}
		pagePtr = &page
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid limit",
			})
			return
		}
		limitPtr = &limit
	}

	addresses, err := h.service.GetAddresses(pagePtr, limitPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}

func (h *Handler) UpdateAddress(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid address id",
		})
		return
	}

	var req UpdateAddressBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedAddress, err := h.service.UpdateAddress(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedAddress)
}

func (h *Handler) DeleteAddress(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid address id",
		})
		return
	}

	err = h.service.DeleteAddress(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "address deleted successfully",
	})
}
