package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type ComponentRepo struct {
	*postgres.Postgres
}

func NewComponentRepo(pg *postgres.Postgres) *ComponentRepo {
	return &ComponentRepo{pg}
}

var _ usecase.ComponentRp = (*ComponentRepo)(nil)

func (c *ComponentRepo) GetComponentsByVendorIDAndTypeID(ctx context.Context, vendorID int64, typeID int64) ([]entity.Component, error) {
	query := `SELECT id, vendor_id, type_id, name, additional_info FROM get_components_by_vendor_and_type($1, $2)`

	rows, err := c.Pool.Query(ctx, query, vendorID, typeID)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var components []entity.Component

	for rows.Next() {
		var cmp entity.Component
		err = rows.Scan(&cmp.ID,
			&cmp.VendorID,
			&cmp.TypeID,
			&cmp.Name,
			&cmp.AdditionalInfo)
		if err != nil {
			return nil, fmt.Errorf("error in parsing component: %w", err)
		}
		components = append(components, cmp)
	}
	return components, nil
}

func (c *ComponentRepo) GetAllComponents(ctx context.Context) ([]entity.Component, error) {
	query := `SELECT * FROM component`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var components []entity.Component

	for rows.Next() {
		var cmp entity.Component
		err = rows.Scan(&cmp.ID,
			&cmp.VendorID,
			&cmp.TypeID,
			&cmp.Name,
			&cmp.AdditionalInfo)
		if err != nil {
			return nil, fmt.Errorf("error in parsing component: %w", err)
		}
		components = append(components, cmp)
	}
	return components, nil
}
