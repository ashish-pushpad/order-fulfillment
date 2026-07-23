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

func (h *Handler) CreateProduct(c *gin.Context) {
	var req CreateProductBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error to Bind Json ", "error ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := h.service.CreateProduct(req)
	if err != nil {
		// log.Println("error to add product ",err)
		slog.Error("error to add the Product ", "error", err)
		c.JSON(500, gin.H{
			"message": "Somethign went wrong ",
			"Product": CreateProductBody{},
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "Product created",
		"Product": req,
	})

}

func (h *Handler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Product id",
		})
		return
	}

	Product, err := h.service.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product fetch successfully",
		"Product": Product,
	})
}

func (h *Handler) GetProducts(c *gin.Context) {
    limitStr := c.Query("limit")
    pageStr := c.Query("page")

    var (
        page  *int
        limit *int
    )

    if pageStr != "" {
        p, err := strconv.Atoi(pageStr)
        if err != nil || p < 1 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
            return
        }
        page = &p
    }

    if limitStr != "" {
        l, err := strconv.Atoi(limitStr)
        if err != nil || l < 1 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
            return
        }
        limit = &l
    }

    products, err := h.service.GetProducts(limit, page)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Productds fetch successfully",
		"Product": products,
	})
}

func (h *Handler) UpdateProduct(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Product id",
		})
		return
	}

	var data UpdateProductBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.Name == "" || data.Price <= 0 || data.Description == "" || data.ImgUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "all fields are required",
		})
		return
	}

	updatedProduct, err := h.service.UpdateProduct(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Product id",
		})
		return
	}

	deletedProduct, err := h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
		"Product": deletedProduct,
	})
}
