package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type SubsidyUseCase struct {
	repo SubsidyRp
}

func NewSubsidyUseCase(r SubsidyRp) *SubsidyUseCase {
	return &SubsidyUseCase{repo: r}
}

func (sb *SubsidyUseCase) GetAllSubsidies(ctx context.Context) ([]entity.Subsidy, error) {
	listSubsidies, err := sb.repo.GetSubsidies(ctx)
	if err != nil {
		return nil, fmt.Errorf("SubsidyUseCase - subsidy list - m.repo.GetSubsidies: %w", err)
	}
	return listSubsidies, nil
}
