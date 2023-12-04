package app

import (
	"D/LatihanAji/GolangBelajar/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RunApp() {
	db := config.ConnectDB()
	config.MigrateModel(db, GetListModel()...)

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
