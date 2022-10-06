package usecase

import (
	"context"
	"pahan/internal/entity"
)

type OrdersUseCase struct {
	repo OrderRp
}

var _ Order = (*OrdersUseCase)(nil)

func NewOrdersUseCase(r OrderRp) *OrdersUseCase {
	return &OrdersUseCase{repo: r}
}

func (o *OrdersUseCase) CreateOrder(ctx context.Context, order entity.Order) error {
	return o.repo.CreateNewOrder(ctx, order)
}

func (o *OrdersUseCase) GetOrders(ctx context.Context) ([]entity.Order, error) {
	return o.repo.GetAllOrders(ctx)
}
