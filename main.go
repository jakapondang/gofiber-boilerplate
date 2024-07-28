package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"gofiber-boilerplatev3/internal/v1/interface/http/routes"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/infra/database"
	"gofiber-boilerplatev3/pkg/msg"
	"gofiber-boilerplatev3/pkg/utils"
)

func main() {
	// Load configurations
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize Logrus
	utils.NewLogger()

	//setup fiber
	app := fiber.New(utils.NewFiberError())
	app.Use(recover.New())
	app.Use(cors.New())

	// Initialize the database connection
	database.Connect(config.AppConfig)

	// Setup routes
	v1.SetupUserRoutes(app, database.DB, config.AppConfig)

	//log.Fatal(app.Listen(":" + config.AppConfig.Server.Port))
	err = app.Listen(":" + config.AppConfig.Server.Port)
	msg.PanicLogging(err)
}
