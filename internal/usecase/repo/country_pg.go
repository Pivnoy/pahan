package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type CountryRepo struct {
	*postgres.Postgres
}

var _ usecase.CountryRp = (*CountryRepo)(nil)

func NewCountryRepo(pg *postgres.Postgres) *CountryRepo {
	return &CountryRepo{pg}
}

func (c *CountryRepo) GetCountries(ctx context.Context) ([]entity.Country, error) {
	query := `SELECT * FROM country`

	rows, err := c.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var countries []entity.Country

	for rows.Next() {
		var cn entity.Country
		err = rows.Scan(&cn.ID, &cn.Name, &cn.GdpUSD)
		if err != nil {
			return nil, fmt.Errorf("cannot parse all fields into struct: %w", err)
		}
		countries = append(countries, cn)
	}
	return countries, nil
}
