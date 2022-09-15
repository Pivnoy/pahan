package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type EngineUseCase struct {
	repo EngineRp
}

func NewEngineUseCase(r EngineRp) *EngineUseCase {
	return &EngineUseCase{
		repo: r,
	}
}

func (uc *EngineUseCase) Engines(ctx context.Context) ([]entity.Engine, error) {
	listEngine, err := uc.repo.GetEngines(ctx)
	if err != nil {
		return nil, fmt.Errorf("EngineUseCase - engine list - s.repo.GetEngines: %w", err)
	}
	return listEngine, nil
}
