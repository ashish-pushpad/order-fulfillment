package inventories


import "time"

type Warehouse struct {
    ProductID   int64 
	WarehouseID int64
	Quantity    int 
	reserved_quantity int
    CreatedAt   time.Time
    UpdatedAt   time.Time
}