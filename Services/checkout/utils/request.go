package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromRequest(c *gin.Context) (int64, error) {
	value, exists := c.Get("user_id")
	if !exists {
		return 0,errors.New("user id not found")
	}
	
	userID, ok := value.(int64)
	if !ok {
		return 0, errors.New("invalid user id")
	}

	return userID, nil
}




type CartItemResponse struct {
	ID        int64
	ProductID int64

	Quantity int

	UnitPrice float64

	TotalPrice float64

	ProductName        string
    ProductDescription string
    ProductImgUrl      string
}


type CartResponse struct {

	ID int64

	UserID int64

	Items []CartItemResponse

	Subtotal float64

	TotalItems int
}