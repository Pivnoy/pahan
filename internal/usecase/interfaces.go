package usecase

import (
	"context"
	"pahan/internal/entity"
)

type (
	Model interface {
		NewModel(context.Context, entity.Model) error
		GetAllModels(context.Context) ([]entity.Model, error)
	}

	ModelRp interface {
		DoNewModel(context.Context, entity.Model) error
		GetModels(context.Context) ([]entity.Model, error)
	}

	CountryRp interface {
		GetCountries(context.Context) ([]entity.Country, error)
	}

	Vendor interface {
		GetAllVendors(context.Context) ([]entity.Vendor, error)
	}

	VendorRp interface {
		GetVendors(context.Context) ([]entity.Vendor, error)
	}

	Order interface {
		NewOrder(context.Context, entity.Order) error
	}

	OrderRp interface {
		DoNewOrder(context.Context, entity.Order) error
	}

	Shipment interface {
		NewShipment(context.Context, entity.Shipment) error
	}

	ShipmentRp interface {
		DoNewShipment(context.Context, entity.Shipment) error
	}
)
