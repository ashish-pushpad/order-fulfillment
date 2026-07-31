package internal


type CreateAddressBody struct{
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	HouseNo string `json:"house_no" binding:"required"`
	Apartment string `json:"apartment" `
	Colony string `json:"colony" binding:"required"`

	City string `json:"city" binding:"required"`
	State string `json:"state" binding:"required"`
	PinCode string `json:"pin_code" binding:"required"`
	IsDefault bool `json:"is_default" binding:"required"`
}

type AddressResponse struct {
	Id int64 	`json:"id" `
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	HouseNo string `json:"house_no" `
	Apartment string `json:"apartment" `
	Colony string `json:"colony" `

	City string `json:"city" `
	State string `json:"state" `

	UserId int64 	`json:"user_id"`
	PinCode string `json:"pin_code" `
	IsDefault bool `json:"is_default" binding:"required"`
}

type GetAddressesResponse struct {
	Addresss []AddressResponse `json:"addresses"`
}

type UpdateAddressBody struct{
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	HouseNo string `json:"house_no"  binding:"required"`
	Apartment string `json:"apartment"  binding:"required"`
	Colony string `json:"colony"  binding:"required"`

	City string `json:"city"  binding:"required"`
	State string `json:"state"  binding:"required"`

	PinCode string `json:"pin_code"  binding:"required"`
	IsDefault bool `json:"is_default" binding:"required"`
}



