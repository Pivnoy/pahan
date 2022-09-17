package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type VendorRepo struct {
	*postgres.Postgres
}

func NewVendorRepo(pg *postgres.Postgres) *VendorRepo {
	return &VendorRepo{pg}
}

var _ usecase.VendorRp = (*VendorRepo)(nil)

func (v *VendorRepo) GetVendors(ctx context.Context) ([]entity.Vendor, error) {
	query := `SELECT * FROM vendor`

	rows, err := v.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var vendors []entity.Vendor

	for rows.Next() {
		var vn entity.Vendor
		err = rows.Scan(&vn.ID, &vn.Name, &vn.CountryID, &vn.Capitalization)
		if err != nil {
			return nil, fmt.Errorf("error in parsing vendor: %w", err)
		}
		vendors = append(vendors, vn)
	}
	return vendors, nil
}
