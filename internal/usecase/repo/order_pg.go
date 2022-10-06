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

func (o *OrderRepo) CreateNewOrder(ctx context.Context, ord entity.Order) error {
	query := `SELECT create_order($1, $2, $3)`

	rows, err := o.Pool.Query(ctx, query, ord.ModelID, ord.Quantity, ord.OrderType)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (o *OrderRepo) GetAllOrders(ctx context.Context) ([]entity.Order, error) {
	query := `SELECT * FROM "order"`

	rows, err := o.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {
		var ord entity.Order
		err = rows.Scan(&ord.ID,
			&ord.ModelID,
			&ord.Quantity,
			&ord.OrderType)
		if err != nil {
			return nil, fmt.Errorf("error in parsing order: %w", err)
		}
		orders = append(orders, ord)
	}
	return orders, nil
}
