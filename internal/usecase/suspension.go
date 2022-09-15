package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type SuspensionUseCase struct {
	repo SuspensionRp
}

func NewSuspensionUseCase(r SuspensionRp) *SuspensionUseCase {
	return &SuspensionUseCase{
		repo: r,
	}
}

func (uc *SuspensionUseCase) Suspensions(ctx context.Context) ([]entity.Suspension, error) {
	listSuspension, err := uc.repo.GetSuspensions(ctx)
	if err != nil {
		return nil, fmt.Errorf("SuspensionUseCase - engine list - s.repo.GetSuspensions: %w", err)

	}
	return listSuspension, nil
}
