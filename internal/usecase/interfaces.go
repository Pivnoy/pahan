package usecase

import (
	"context"
	"pahan/internal/entity"
)

type (
	Engineer interface {
		GetAllEngineerByIdVendor(context.Context, int64) ([]entity.Engineer, error)
	}

	EngineerRp interface {
		GetEngineerByIdVendor(context.Context, int64) ([]entity.Engineer, error)
	}

	Subsidy interface {
		GetAllSubsidies(context.Context) ([]entity.Subsidy, error)
		CreateSubsidy(context.Context, int64, float64, string) error
	}

	SubsidyRp interface {
		GetSubsidies(context.Context) ([]entity.Subsidy, error)
		CreateAndLinkSubsidy(context.Context, int64, float64, string) error
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

	FactoryRp interface {
		GetFactoriesByVendorID(context.Context, int64) ([]entity.Factory, error)
	}

	Factory interface {
		GetFactoriesByVendor(context.Context, int64) ([]entity.Factory, error)
	}

	ComponentRp interface {
		GetComponentsByVendorIDAndTypeID(context.Context, int64, int64) ([]entity.Component, error)
	}

	Component interface {
		GetComponentsByVendorAndType(context.Context, int64, int64) ([]entity.Component, error)
	}

	TypeRp interface {
		GetAllTypes(context.Context) ([]entity.Type, error)
	}

	Type interface {
		GetTypes(context.Context) ([]entity.Type, error)
	}
)
