package entity

type Subsidy struct {
	ID           int64 `json:"id"`
	CountryID    int64 `json:"country_id"`
	RequirePrice int64 `json:"require_price"`
	RequiredWd   int64 `json:"required_wd"`
}
