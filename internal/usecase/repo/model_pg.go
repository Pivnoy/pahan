package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type ModelRepo struct {
	*postgres.Postgres
}

var _ usecase.ModelRp = (*ModelRepo)(nil)

func NewModelRepo(pg *postgres.Postgres) *ModelRepo {
	return &ModelRepo{pg}
}

// DoNewModel Создание новой модели
func (m *ModelRepo) DoNewModel(ctx context.Context, car entity.Model) error {
	query := `SELECT do_new_model($1, $2, $3, $4, $5, $6, $7)`

	rows, err := m.Pool.Query(ctx, query, car.WheelDrive, car.Significance, car.ProdCost, car.ProdCost, car.EngineID, car.SuspensionID, car.Name)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
