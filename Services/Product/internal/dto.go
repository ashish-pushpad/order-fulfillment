package internal


type CreateProductBody struct{
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	ImgUrl string `json:"img_url" `
}




type GetProductResponse struct {
	Id int64 	`json:"id"`
	Name string `json:"name" `
	Description string `json:"description" `
	Price float64 `json:"price" `
	ImgUrl string `json:"img_url" `
}

type GetProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}

type UpdateProductBody struct{
	Name string `json:"name" binding:"required"`
	Description string `json:"description"  binding:"required"`
	Price float64 `json:"price" binding:"required"`
	ImgUrl string `json:"img_url" binding:"required"`
}

type  UpdateProductBodyRes struct{
	Id int64 	`json:"id" `
	Name string `json:"name" `
	Description string `json:"description" `
	Price float64 `json:"price" `
	ImgUrl string `json:"img_url" `
}


type GetProductsByIdsResponse struct {
	Id int64 	`json:"id"`
	Name string `json:"name" `
	Description string `json:"description" `
	ImgUrl string `json:"img_url" `
}


type GetProductPriceByIdResponse struct {
		Id int64  `json:"id"`
		Price float64 `json:"price"`
	}