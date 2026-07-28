package order

import "time"

type Order struct {
	ID int64 `json:"id"`

	UserID      int64  `json:"user_id"`
	AddressID   int64  `json:"address_id"`
	// WarehouseID *int64 `json:"warehouse_id,omitempty"`

	OrderNumber string `json:"order_number"`

	Status string `json:"status"`

	Subtotal     float64 `json:"subtotal"`
	ShippingCost float64 `json:"shipping_cost"`
	Tax          float64 `json:"tax"`
	Discount     float64 `json:"discount"`
	TotalAmount  float64 `json:"total_amount"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID int64 `json:"id"`

	OrderID int64 `json:"order_id"`
 
	ProductID int64 `json:"product_id"`
    WarehouseID *int64 `json:"warehouse_id,omitempty"`
	Quantity int `json:"quantity"`

	UnitPrice float64 `json:"unit_price"`

	TotalPrice float64 `json:"total_price"`
}