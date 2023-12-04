package app

import (
	"D/LatihanAji/GolangBelajar/router"

	"github.com/gofiber/fiber/v2"
)

type SetupRouterStruct struct {
	BookRouter router.BookRouter
}

func setupRouter(setupUseCase *SetupUseCaseStruct) SetupRouterStruct {
	return SetupRouterStruct{
		BookRouter: router.NewBookRouter(&setupUseCase.BookUseCase),
	}
}

func setupRouting(r SetupRouterStruct, app *fiber.App) {
	r.BookRouter.Route(app)
}
