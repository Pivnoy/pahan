package entity

type Vendor struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	CountryID      int64  `json:"country_id"`
	Capitalization int64  `json:"capitalization"`
}
