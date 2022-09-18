package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type ShipmentUseCase struct {
	repo ShipmentRp
}

func NewShipmentUseCase(s ShipmentRp) *ShipmentUseCase {
	return &ShipmentUseCase{repo: s}
}

var _ Shipment = (*ShipmentUseCase)(nil)

func (s ShipmentUseCase) NewShipment(ctx context.Context, shipment entity.Shipment) error {
	err := s.repo.DoNewShipment(ctx, shipment)
	if err != nil {
		return fmt.Errorf("ShipmentUseCase - NewShipment - s.repo.DoNewShipment: %w", err)
	}
	return nil
}
