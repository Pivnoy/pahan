package entity

type Model struct {
	ID           int64   `json:"id"`
	WheelDrive   string  `json:"wheeldrive"`
	Significance int64   `json:"significance"`
	Price        int64   `json:"price"`
	ProdCost     float64 `json:"prod_cost"`
	EngineID     int64   `json:"engine_id"`
	SuspensionID int64   `json:"suspension_id"`
	VendorID     int64   `json:"vendor_id"`
	Name         string  `json:"name"`
}
