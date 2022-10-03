package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type SubsidyRepo struct {
	*postgres.Postgres
}

func NewSubsidyRepo(pg *postgres.Postgres) *SubsidyRepo {
	return &SubsidyRepo{pg}
}

var _ usecase.SubsidyRp = (*SubsidyRepo)(nil)

func (sr *SubsidyRepo) GetSubsidies(ctx context.Context) ([]entity.Subsidy, error) {
	query := `SELECT id, country_id, require_price, required_wd FROM get_all_subsidies()`

	rows, err := sr.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var subsidies []entity.Subsidy

	for rows.Next() {
		var subsidy entity.Subsidy
		err = rows.Scan(
			&subsidy.ID,
			&subsidy.CountryID,
			&subsidy.RequirePrice,
			&subsidy.RequiredWd,
		)

		if err != nil {
			return nil, fmt.Errorf("error in parsing subsidy: %w", err)
		}

		subsidies = append(subsidies, subsidy)
	}
	return subsidies, nil
}

func (sr *SubsidyRepo) CreateAndLinkSubsidy(ctx context.Context, countryIDBy int64, requirePriceBy float64, requiredWdBy string) error {
	query := `create_subsidy($1, $2, $3)`

	_, err := sr.Pool.Query(ctx, query, countryIDBy, requirePriceBy, requiredWdBy)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	return nil
}