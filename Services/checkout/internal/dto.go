package internal

type CheckoutBody struct {
	AddressID int64 `json:"address_id" binding:"required"`
}

type CheckoutResponse struct {
	OrderID     int64   `json:"order_id"`
	OrderNumber string  `json:"order_number"`
	Subtotal    float64 `json:"subtotal"`
	Shipping    float64 `json:"shipping"`
	Tax         float64 `json:"tax"`
	Discount    float64 `json:"discount"`
	Total        float64 `json:"total"`
	Status      string  `json:"status"`
}