package controllers

import (
	"strconv"

	"github.com/HenCor2019/task-go/api/common/responses"
	"github.com/gofiber/fiber/v2"
)

func (shipengineController *Controller) EstimateManyCarriers(c *fiber.Ctx) error {
	quantityEstimations, err := strconv.Atoi(c.Query("qty"))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": "Error in quantityEstimations",
		})
	}

	pokemons := shipengineController.service.FetchManyEstimations(quantityEstimations)
	return common.SuccessResponse(c, pokemons, fiber.StatusOK)
}
