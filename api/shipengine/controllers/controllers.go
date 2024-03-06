package controllers

import (
	"github.com/HenCor2019/task-go/api/shipengine/services"
	"github.com/gofiber/fiber/v2"
)

type ShipengineController interface {
	EstimateManyCarriers(c *fiber.Ctx) error
}

type Controller struct {
	service services.ShipengineService
}

func New(service services.ShipengineService) ShipengineController {
	return &Controller{service}
}
