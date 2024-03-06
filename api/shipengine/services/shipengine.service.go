package services

import (
	"github.com/HenCor2019/task-go/api/models"
	"github.com/gofiber/fiber/v2"
)

func (shipengineService *Service) FetchManyEstimations(quantityEstimations int) []models.Shipengine {
	estimations, err := shipengineService.gateway.FetchRatesEstimation(quantityEstimations)
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	return estimations
}
