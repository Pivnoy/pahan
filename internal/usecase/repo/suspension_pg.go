package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type SuspensionRepo struct {
	*postgres.Postgres
}

func NewSuspensionRepo(pg *postgres.Postgres) *SuspensionRepo {
	return &SuspensionRepo{pg}
}

var _ usecase.SuspensionRp = (*SuspensionRepo)(nil)

func (s *SuspensionRepo) GetSuspensions(ctx context.Context) ([]entity.Suspension, error) {
	query := `SELECT * FROM suspension`

	rows, err := s.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var ssp []entity.Suspension

	for rows.Next() {
		var sp entity.Suspension
		err = rows.Scan(&sp.ID, &sp.Name, &sp.Type)
		if err != nil {
			return nil, fmt.Errorf("cannot parse all fields into struct: %w", err)
		}
		ssp = append(ssp, sp)
	}
	return ssp, nil
}
