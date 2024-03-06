package api

import (
	"fmt"
	"os"

	"github.com/spf13/cast"

	"github.com/HenCor2019/task-go/api/common/responses"
	"github.com/HenCor2019/task-go/api/shipengine/controllers"

	"github.com/HenCor2019/task-go/api/pokemons/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	PokemonControllers    PokemonsControllers.PokemonController
	ShipengineControllers controllers.ShipengineController
}

func New(
	pokemonControllers PokemonsControllers.PokemonController,
	shipengineControllers controllers.ShipengineController,
) *API {
	return &API{
		PokemonControllers:    pokemonControllers,
		ShipengineControllers: shipengineControllers,
	}
}

func (api *API) Start(app *fiber.App) error {
	app.Use(recover.New())
	v1 := app.Group("api/v1")

	v1.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	v1.Get("pokemons/", api.PokemonControllers.FindManyById)
	v1.Post("shipengine/", api.ShipengineControllers.EstimateManyCarriers)

	v1.Use(common.NotFoundHandler)
	PORT := os.Getenv("PORT")
	return app.Listen(fmt.Sprintf(":%s", cast.ToString(PORT)))
}
