package internal


import "time"

type Address struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`

	HouseNo     string    `json:"house_no"`
	Apartment   string    `json:"apartment"`
	Colony      string    `json:"colony"`

	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	PinCode     string    `json:"pin_code"`

	IsDefault   bool      `json:"is_default"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}