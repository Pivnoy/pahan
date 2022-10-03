package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type OrderRepo struct {
	*postgres.Postgres
}

var _ usecase.OrderRp = (*OrderRepo)(nil)

func NewOrdersRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}

func (o OrderRepo) DoNewOrder(ctx context.Context, order entity.Order) error {
	query := `SELECT do_new_order($1, $2, $3)`

	_, err := o.Pool.Query(ctx, query, order.ModelID, order.Quantity, order.OrderType)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	return nil
}
