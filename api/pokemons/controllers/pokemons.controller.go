package PokemonsControllers

import (
	"strconv"

	"github.com/HenCor2019/task-go/api/common/responses"
	"github.com/HenCor2019/task-go/api/pokemons/dtos"
	"github.com/gofiber/fiber/v2"
)

func (pokemonController *Controller) FindManyById(c *fiber.Ctx) error {
	pokemonId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"content": nil,
			"message": "Please insert a valid pokemon id",
		})
	}

	var fetchPokemonIds = pokemonsDtos.FetchPokemonsByIdsDto{
		Ids: []int{pokemonId},
	}
	pokemons := pokemonController.service.FindManyById(fetchPokemonIds)
	return common.SuccessResponse(c, pokemons, fiber.StatusOK)
}
