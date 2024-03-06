package main

import (
	"context"

	"github.com/HenCor2019/task-go/api"

	"github.com/HenCor2019/task-go/api/pokemons/controllers"
	"github.com/HenCor2019/task-go/api/pokemons/gateways"
	"github.com/HenCor2019/task-go/api/pokemons/services"

	ShipengineControllers "github.com/HenCor2019/task-go/api/shipengine/controllers"
	ShipengineGateways "github.com/HenCor2019/task-go/api/shipengine/gateways"
	ShipengineServices "github.com/HenCor2019/task-go/api/shipengine/services"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			PokemonsGateways.New,
			PokemonsServices.New,
			PokemonsControllers.New,
		),

		fx.Provide(
			ShipengineGateways.New,
			ShipengineServices.New,
			ShipengineControllers.New,
		),

		fx.Provide(
			api.New,
			fiber.New,
		),

		fx.Invoke(setLifeCycle),
	)

	app.Run()
}

func setLifeCycle(
	lc fx.Lifecycle,
	a *api.API,
	app *fiber.App,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go a.Start(app) // nolint

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
