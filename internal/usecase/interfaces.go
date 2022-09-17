package usecase

import (
	"context"
	"pahan/internal/entity"
)

type (
	Model interface {
		// проектирование новой модели
		NewModel(context.Context, entity.Model) error
	}

	ModelRp interface {
		//TODO
		DoNewModel(context.Context, entity.Model) error
	}

	CountryRp interface {
		GetCountries(context.Context) ([]entity.Country, error)
	}

	VendorRp interface {
		GetVendors(context.Context) ([]entity.Vendor, error)
	}

	Suspension interface {
		Suspensions(context.Context) ([]entity.Suspension, error)
	}

	SuspensionRp interface {
		GetSuspensions(context.Context) ([]entity.Suspension, error)
	}

	Engine interface {
		Engines(context.Context) ([]entity.Engine, error)
	}

	EngineRp interface {
		GetEngines(context.Context) ([]entity.Engine, error)
	}
)
