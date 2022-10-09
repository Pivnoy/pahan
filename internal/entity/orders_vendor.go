package entity

type OrdersVendor struct {
	ModelName string `json:"model_name"`
	ModelID   int64  `json:"model_id"`
	OrderID   int64  `json:"order_id"`
	Quantity  int64  `json:"quantity"`
	OrderType string `json:"order_type"`
}
