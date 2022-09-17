package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type ShipmentRepo struct {
	*postgres.Postgres
}

func NewShipmentRepo(pg *postgres.Postgres) *ShipmentRepo {
	return &ShipmentRepo{pg}
}

var _ usecase.ShipmentRp = (*ShipmentRepo)(nil)

func (s ShipmentRepo) DoNewShipment(ctx context.Context, shipment entity.Shipment) error {
	query := `SELECT do_new_shipment($1, $2, $3)`

	rows, err := s.Pool.Query(ctx, query, shipment.OrderID, shipment.CountryToID, shipment.Date)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
