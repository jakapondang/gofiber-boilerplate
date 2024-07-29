package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/internal/v1/interface/http/routes"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/infra/database"
	"gofiber-boilerplatev3/pkg/utils"
	"gofiber-boilerplatev3/pkg/utils/logruspack"
)

func main() {
	// Initialize Logrus
	logruspack.Init()

	// Load configurations
	config.Init()

	//setup fiber
	app := fiber.New(utils.NewFiberError())
	app.Use(recover.New())
	app.Use(cors.New())

	// Register the request ID and logging middleware
	app.Use(middlewares.RequestID())
	app.Use(middlewares.LogRequestResponse())

	// Initialize the database connection
	if err := database.Connect(config.AppConfig); err != nil {
		logruspack.Logger.Fatalf("Failed to connect to the database: %v", err)
	}

	// Setup routes
	v1.SetupUserRoutes(app, database.DB, config.AppConfig)

	//log.Fatal(app.Listen(":" + config.AppConfig.Server.Port))
	err := app.Listen(":" + config.AppConfig.Server.Port)
	if err != nil {
		logruspack.Logger.Fatalf("Failed to connect to the port : %v", err)
	}
}
