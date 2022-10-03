package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type SubsidyUseCase struct {
	repo SubsidyRp
}

var _ Subsidy = (*SubsidyUseCase)(nil)

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

func (sb *SubsidyUseCase) CreateSubsidy(ctx context.Context, countryIDBy int64, requirePriceBy float64, requiredWdBy string) error {
	return sb.repo.CreateAndLinkSubsidy(ctx, countryIDBy, requirePriceBy, requiredWdBy)
}
