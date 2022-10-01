package usecase

import (
	"context"
	"pahan/internal/entity"
)

type (

	Subsidy interface {
		GetAllSubsidies(context.Context) ([]entity.Subsidy, error)
	}

	SubsidyRp interface {
		GetSubsidies(context.Context) ([]entity.Subsidy, error)
	}


	Model interface {
		NewModel(context.Context, entity.Model) error
		GetAllModels(context.Context) ([]entity.Model, error)
	}

	ModelRp interface {
		DoNewModel(context.Context, entity.Model) error
		GetModels(context.Context) ([]entity.Model, error)
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
