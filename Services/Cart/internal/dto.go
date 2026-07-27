package internal


type AddCartItemBody struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int   `json:"quantity" binding:"required"`
}

type UpdateCartItemBody struct {
	Quantity int `json:"quantity" binding:"required"`
}

type CartItemResponse struct {
	ID        int64
	ProductID int64

	Quantity int

	UnitPrice float64

	TotalPrice float64

	ProductName        string
    ProductDescription string
    ProducImgUrl      string
}

type CartResponse struct {

	ID int64

	UserID int64

	Items []CartItemResponse

	Subtotal float64

	TotalItems int
}