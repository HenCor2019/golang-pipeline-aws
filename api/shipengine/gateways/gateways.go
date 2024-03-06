package gateways

import (
	"github.com/HenCor2019/task-go/api/models"
)

type ShipengineGateway interface {
	FetchRatesEstimation(quantityEstimations int) ([]models.Shipengine, error)
}

type Gateway struct{}

func New() ShipengineGateway {
	return &Gateway{}
}
