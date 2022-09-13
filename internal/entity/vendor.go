package entity

type Vendor struct {
	ID             int64   `json:"id"`
	CountryID      int64   `json:"country_id"`
	Capitalization float64 `json:"capitalization"`
	Name           string  `json:"name"`
}
