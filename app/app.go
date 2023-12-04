package app

import (
	"D/LatihanAji/GolangBelajar/config"
	"D/LatihanAji/GolangBelajar/model"
	"D/LatihanAji/GolangBelajar/repository"
	"D/LatihanAji/GolangBelajar/router"
	"D/LatihanAji/GolangBelajar/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RunApp() {
	db := config.ConnectDB()
	config.MigrateModel(db, &model.Book{})

	// setup repo
	repo := setupRepository(db)

	// setup usecase
	uc := setupUseCase(&repo, db)

	// setup roter
	r := setupRouter(&uc)

	// start server
	app := fiber.New()

	// setup routing
	setupRouting(r, app)

	err := app.Listen(":3400")
	if err != nil {
		log.Fatal(err)
	}
}

type SetupRepositoryStruct struct {
	BookRepository repository.BookRepository
}

func setupRepository(database *gorm.DB) SetupRepositoryStruct {
	return SetupRepositoryStruct{
		BookRepository: repository.NewBookRepository(database),
	}
}

type SetupUseCaseStruct struct {
	BookUseCase usecase.BookUseCase
}

func setupUseCase(setupRepo *SetupRepositoryStruct, db *gorm.DB) SetupUseCaseStruct {
	return SetupUseCaseStruct{
		BookUseCase: usecase.NewBookUseCase(setupRepo.BookRepository, db),
	}
}

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
