package usecase

import (
	"context"
	"pahan/internal/entity"
)

type (
	ModelRp interface {
		//TODO
	}

	CountryRp interface {
		GetCountries(context.Context) ([]entity.Country, error)
	}

	VendorRp interface {
		GetVendors(context.Context) ([]entity.Vendor, error)
	}

	SuspensionRp interface {
		GetSuspensions(context.Context) ([]entity.Suspension, error)
	}

	EngineRp interface {
		GetEngines(context.Context) ([]entity.Engine, error)
	}
)
