package internal


import "time"

type Warehouse struct {
    ID   int64
    Name string 
	City string 
	Address string 
    CreatedAt   time.Time
    UpdatedAt   time.Time
}