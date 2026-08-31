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

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Error to Bind Json ", "error ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := h.service.CreateUser(c,req)
	if err != nil {
		slog.Error("error to add the user ", "error", err)
		c.JSON(500, gin.H{
			"message": "Somethign went wrong ",
			"user":    CreateUserBody{},
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "User created",
		"user":    req,
	})

}

func (h *Handler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	user, err := h.service.GetUser(ctx,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user fetch successfully",
		"user":    user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var data UpdateUserBody
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.Name == "" || data.Email == "" || data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "all fields are required",
		})
		return
	}

	updatedUser, err := h.service.UpdateUser(c ,id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	deletedUser, err := h.service.DeleteUser(c,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"user":    deletedUser,
	})
}

func (h *Handler) LoginUser(c *gin.Context) {
	var loginDetails GetLoginRequest
	err := c.ShouldBindJSON(&loginDetails)
	if err != nil {
		slog.Error("Error to Bind Json ", "error ", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// authToken, err := h.service.LoginUser(loginDetails.Email, loginDetails.Password)
	var authToken string
	// if err != nil {
	// 	slog.Error("Error to Login user ", "error ", err)
	// 	c.JSON(500, gin.H{
	// 		"message": err.Error(),
	// 	})
	// }
	c.SetCookie(
		"access_token", // Cookie name
		authToken,      // Value
		3600*24,        // MaxAge (24 hours)
		"/",            // Path
		"",             // Domain (empty = current domain)
		false,          // Secure (true in production with HTTPS)
		true,           // HttpOnly
	)
	c.JSON(200, gin.H{
		"message": "Loginn successfully",
		"token":   authToken,
	})
}
