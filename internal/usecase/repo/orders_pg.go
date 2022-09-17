package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type OrdersRepo struct {
	*postgres.Postgres
}

var _ usecase.OrdersRp = (*OrdersRepo)(nil)

func NewOrdersRepo(pg *postgres.Postgres) *OrdersRepo {
	return &OrdersRepo{pg}
}

func (o OrdersRepo) DoNewOrder(ctx context.Context, order entity.Orders) error {
	query := `SELECT do_new_order($1, $2, $3)`

	rows, err := o.Pool.Query(ctx, query, order.ModelID, order.Quantity, order.OrderType)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
