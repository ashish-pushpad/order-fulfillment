package internal


type CreateWarehouseBody struct{
	Name string `json:"name" binding:"required"`
	City string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
}


type GetWarehouseResponse struct {
	Id int64 	`json:"id"`
	Name string `json:"name" binding:"required"`
	City string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type GetWarehousesResponse struct {
	Warehouses []GetWarehouseResponse `json:"products"`
}

type UpdateWarehouseBody struct{
	Name string `json:"name" binding:"required"`
	City string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateWarehouseBodyRes struct{
	Id int64 	
	Name string 
	City string 
	Address string
}

