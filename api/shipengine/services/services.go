package services

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/HenCor2019/task-go/api/shipengine/gateways"
)

type ShipengineService interface {
	FetchManyEstimations(quantityEstimations int) []models.Shipengine
}

type Service struct {
	gateway gateways.ShipengineGateway
}

func New(gateway gateways.ShipengineGateway) ShipengineService {
	return &Service{gateway}
}
