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

func (o *OrderRepo) GetAllOrdersByVendor(ctx context.Context, vendorID int64) ([]entity.OrdersVendor, error) {
	query := `select model_name, model_id, order_id, quantity, order_type from get_orders_by_vendor_id(1)`

	rows, err := o.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var ordersVendor []entity.OrdersVendor

	for rows.Next() {
		var orderVendor entity.OrdersVendor
		err = rows.Scan(&orderVendor.ModelName,
			&orderVendor.ModelID,
			&orderVendor.OrderID,
			&orderVendor.Quantity,
			&orderVendor.OrderType)
		if err != nil {
			return nil, fmt.Errorf("error in parsing order: %w", err)
		}
		ordersVendor = append(ordersVendor, orderVendor)
	}
	return ordersVendor, nil
}
