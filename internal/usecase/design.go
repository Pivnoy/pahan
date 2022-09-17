package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

// Сконструировать новую модель

type DesignUseCase struct {
	repo ModelRp
}

func NewDesignUseCase(r ModelRp) *DesignUseCase {
	return &DesignUseCase{
		repo: r,
	}
}

func (uc *DesignUseCase) NewModel(ctx context.Context, car entity.Model) error {
	err := uc.repo.DoNewModel(ctx, car)
	if err != nil {
		return fmt.Errorf("DesignUseCase - NewModel - s.repo.DoNewModel: %w", err)
	}
	return nil
}
