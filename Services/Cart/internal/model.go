package internal


import "time"



type CartTime struct {
	CartID    int64     `json:"cart_id"`
	ExpiresAt time.Time `json:"expires_at"`
}