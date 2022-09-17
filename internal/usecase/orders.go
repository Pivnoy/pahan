package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type OrdersUseCase struct {
	repo OrdersRp
}

var _ Orders = (*OrdersUseCase)(nil)

func NewOrdersUseCase(r OrdersRp) *OrdersUseCase {
	return &OrdersUseCase{repo: r}
}

func (o OrdersUseCase) NewOrder(ctx context.Context, order entity.Orders) error {
	err := o.repo.DoNewOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("OrdersUseCase - NewOrder - s.repo.DoNewOrder: %w", err)
	}
	return nil
}
