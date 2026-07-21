package internal


type CreateInventoryBody struct {
	ProductID   int64 `json:"product_id" binding:"required"`
	WarehouseID int64 `json:"warehouse_id" binding:"required"`
	Quantity    int   `json:"quantity" binding:"required"`
}


type InventoryResponse struct {
	ID               int64 `json:"id"`
	ProductID        int64 `json:"product_id"`
	WarehouseID      int64 `json:"warehouse_id"`
	Quantity         int   `json:"quantity"`
	ReservedQuantity int   `json:"reserved_quantity"`
	AvailableQuantity int 	`jons:"reserved_quantity"`
}


type UpdateInventoryBody struct {
	Quantity int `json:"quantity" binding:"required"`
}


type GetInventoriesResponse struct {
	Inventories []InventoryResponse `json:"inventories"`
}