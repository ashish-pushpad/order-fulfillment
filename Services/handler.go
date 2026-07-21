package inventories

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

// ----------------------------
// CREATE
// ----------------------------

func (h *Handler) CreateInventory(c *gin.Context) {
	var req CreateInventoryBody

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind json", "error", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.CreateInventory(req)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "inventory created successfully",
	})
}

// ----------------------------
// GET ONE
// ----------------------------

func (h *Handler) GetInventory(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inventory id",
		})
		return
	}

	inventory, err := h.service.GetInventory(id)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// ----------------------------
// GET ALL
// ----------------------------

func (h *Handler) GetInventories(c *gin.Context) {

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

	inventories, err := h.service.GetInventories(pagePtr, limitPtr)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"inventories": inventories,
	})
}

// ----------------------------
// UPDATE
// ----------------------------

func (h *Handler) UpdateInventory(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inventory id",
		})
		return
	}

	var req UpdateInventoryBody

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	inventory, err := h.service.UpdateInventory(id, req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// ----------------------------
// DELETE
// ----------------------------

func (h *Handler) DeleteInventory(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid inventory id",
		})

		return
	}

	err = h.service.DeleteInventory(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "inventory deleted successfully",
	})
}

// ----------------------------
// RESERVE STOCK
// ----------------------------

// func (h *Handler) ReserveStockTx(c *gin.Context) {

// 	productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
// 	warehouseID, _ := strconv.ParseInt(c.Param("warehouseId"), 10, 64)

// 	qty, err := strconv.Atoi(c.Param("qty"))

// 	if err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid quantity",
// 		})

// 		return
// 	}

// 	err = h.service.ReserveStockTx(productID, warehouseID, qty)

// 	if err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "stock reserved",
// 	})
// }

// ----------------------------
// RELEASE STOCK
// ----------------------------

// func (h *Handler) ReleaseStock(c *gin.Context) {

// 	productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
// 	warehouseID, _ := strconv.ParseInt(c.Param("warehouseId"), 10, 64)

// 	qty, err := strconv.Atoi(c.Param("qty"))

// 	if err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid quantity",
// 		})

// 		return
// 	}

// 	err = h.service.ReleaseStock(productID, warehouseID, qty)

// 	if err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "stock released",
// 	})
// }

// ----------------------------
// COMMIT STOCK
// ----------------------------

// func (h *Handler) CommitReservation(c *gin.Context) {

// 	productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
// 	warehouseID, _ := strconv.ParseInt(c.Param("warehouseId"), 10, 64)

// 	qty, err := strconv.Atoi(c.Param("qty"))

// 	if err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid quantity",
// 		})

// 		return
// 	}

// 	err = h.service.CommitReservation(productID, warehouseID, qty)

// 	if err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "inventory updated",
// 	})
// }

// ----------------------------
// AVAILABLE STOCK
// ----------------------------

func (h *Handler) GetAvailableStock(c *gin.Context) {

	productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)
	warehouseID, _ := strconv.ParseInt(c.Param("warehouseId"), 10, 64)

	stock, err := h.service.GetAvailableStock(productID, warehouseID)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"available_stock": stock,
	})
}