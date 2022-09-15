package repo

import (
	"context"
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
func (m *ModelRepo) DoNewModel(ctx context.Context, model entity.Model) error {
	//TODO
	return nil
}
