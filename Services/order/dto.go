package order

import "time"

type CreateOrderRequest struct {
	AddressID int64 `json:"address_id" binding:"required"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type OrderItemResponse struct {
	ProductID   int64   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

type OrderResponse struct {
	ID int64 `json:"id"`

	OrderNumber string `json:"order_number"`

	Status string `json:"status"`

	UserID int64 `json:"user_id"`

	AddressID int64 `json:"address_id"`

	Subtotal float64 `json:"subtotal"`

	ShippingCost float64 `json:"shipping_cost"`

	Tax float64 `json:"tax"`

	Discount float64 `json:"discount"`

	TotalAmount float64 `json:"total_amount"`

	Items []OrderItemResponse `json:"items"`

	CreatedAt time.Time `json:"created_at"`
}
