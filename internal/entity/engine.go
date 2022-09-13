package entity

type Engine struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Power  float64 `json:"power"`
	Torque float64 `json:"torque"`
	Layout string  `json:"layout"`
}
