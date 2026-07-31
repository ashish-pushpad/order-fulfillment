package internal

import (
	"log/slog"
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

func (h *Handler) CreateWarehouse(c *gin.Context) {

	var req CreateWarehouseBody

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error to Bind Json ", "error ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := h.service.CreateWarehouse(req)
	if err != nil {
		slog.Error("error to add the Warehouse ", "error", err)
		c.JSON(500, gin.H{
			"message": "Somethign went wrong ",
			"Warehouse": CreateWarehouseBody{},
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "Warehouse created",
		"Warehouse": req,
	})

}

func (h *Handler) GetWarehouse(c *gin.Context) {
	idParam := c.Param("id")
	ctx:=c.Request.Context()
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Warehouse id",
		})
		return
	}

	Warehouse, err := h.service.GetWarehouse(ctx,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warehouse fetch successfully",
		"Warehouse": Warehouse,
	})
}

func (h *Handler) GetWarehouses(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	var (
		page  *int
		limit *int
		err   error
	)

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil || p < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
			return
		}
		page=&p
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		limit=&l
	}

	Warehouse, err := h.service.GetWarehouses(limit , page)
	if err != nil {
		slog.Error("error to add the Warehouse ", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warehouse fetch successfully",
		"Warehouse": Warehouse,
	})
}

func (h *Handler) UpdateWarehouse(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Warehouse id",
		})
		return
	}

	var data UpdateWarehouseBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.Name == "" || data.City =="" || data.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "all fields are required",
		})
		return
	}

	updatedWarehouse, err := h.service.UpdateWarehouse(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedWarehouse)
}

func (h *Handler) DeleteWarehouse(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Warehouse id",
		})
		return
	}

	deletedWarehouse, err := h.service.DeleteWarehouse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Warehouse deleted successfully",
		"Warehouse": deletedWarehouse,
	})
}
