package entity

import "time"

type Orders struct {
	ID           int64     `json:"id"`
	ModelID      int64     `json:"model_id"`
	Quantity     int64     `json:"quantity"`
	OrderType    string    `json:"order_type"`
	OrderTime    time.Time `json:"order_time"`
	ShippingDate time.Time `json:"shipping_date"`
}
