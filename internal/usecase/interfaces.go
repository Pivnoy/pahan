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
		AcceptSubsidyUs(context.Context, int64, entity.Model, int64, int64, int64, int64) error
	}

	SubsidyRp interface {
		GetSubsidies(context.Context) ([]entity.Subsidy, error)
		CreateAndLinkSubsidy(context.Context, int64, float64, string) error
		AcceptSubsidy(context.Context, int64, entity.Model, int64, int64, int64, int64) error
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
		CreateOrder(context.Context, entity.Order) error
	}

	OrderRp interface {
		CreateNewOrder(context.Context, entity.Order) error
	}

	Shipment interface {
		CreateShipment(context.Context, entity.Shipment) error
	}

	ShipmentRp interface {
		CreateNewShipment(context.Context, entity.Shipment) error
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
