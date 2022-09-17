package usecase

import (
	"context"
	"pahan/internal/entity"
)

type (
	Model interface {
		NewModel(context.Context, entity.Model) error
	}

	ModelRp interface {
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

	Orders interface {
		NewOrder(context.Context, entity.Orders) error
	}

	OrdersRp interface {
		DoNewOrder(context.Context, entity.Orders) error
	}

	Shipment interface {
		NewShipment(context.Context, entity.Shipment) error
	}

	ShipmentRp interface {
		DoNewShipment(context.Context, entity.Shipment) error
	}
)
