package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type EngineRepo struct {
	*postgres.Postgres
}

var _ usecase.EngineRp = (*EngineRepo)(nil)

func NewEngineRepo(pg *postgres.Postgres) *EngineRepo {
	return &EngineRepo{pg}
}

func (e *EngineRepo) GetEngines(ctx context.Context) ([]entity.Engine, error) {
	query := `SELECT * FROM engine`

	rows, err := e.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var eng []entity.Engine

	for rows.Next() {
		var engine entity.Engine
		err = rows.Scan(&engine.ID, &engine.Name, &engine.Power, &engine.Torque, &engine.Layout)
		if err != nil {
			return nil, fmt.Errorf("cannot parse all fields into struct: %w", err)
		}
		eng = append(eng, engine)
	}
	return eng, nil
}
