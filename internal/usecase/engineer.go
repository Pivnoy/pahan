package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type EngineerUseCase struct {
	repo EngineerRp
}

func NewEngineerUseCase(r EngineerRp) *EngineerUseCase {
	return &EngineerUseCase{repo: r}
}

func (en *EngineerUseCase) GetAllEngineerByIdVendor(ctx context.Context, vendor_id int64) ([]entity.Engineer, error) {
	listEngineers, err := en.repo.GetEngineerByIdVendor(ctx, vendor_id)
	if err != nil {
		return nil, fmt.Errorf("EngineerUseCase - Engineer list - m.repo.GetAllEngineerByIdVendor: %w", err)
	}
	return listEngineers, nil
}
