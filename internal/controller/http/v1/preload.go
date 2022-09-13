package v1

import "pahan/internal/usecase"

type preload struct {
	vn usecase.VendorRp
	sp usecase.SuspensionRp
	en usecase.EngineRp
	cn usecase.CountryRp
}
