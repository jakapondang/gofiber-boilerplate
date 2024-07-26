package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"goamartha/common"
	"goamartha/config"
	"goamartha/route"
)

func main() {
	//setup configuration
	config := configuration.New()

	// Initialize Logrus
	common.NewLogger()

	// Connect to the database
	db := configuration.NewDB(config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	route.Borrower(app, db)
	route.Payment(app, db)

	//start app
	common.Logger.Fatal(app.Listen(config.Get("PORT")))
}
