package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	v1 "gofiber-boilerplatev3/internal/v1/interface/http/routes"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/infra/database"
	"gofiber-boilerplatev3/pkg/utils"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
	"gofiber-boilerplatev3/pkg/utils/logruspack"
	"gofiber-boilerplatev3/pkg/utils/mailpack"
	"strconv"
)

func main() {
	// Initialize Logrus
	logruspack.Init()

	// Load configurations
	config.Init()

	// Initialize JWT settings
	// Set JWT config
	jwt.SetConfig(config.AppConfig.JWT)
	fmt.Println(config.AppConfig.Mail)
	mailpack.SetConfig(config.AppConfig.Mail)

	//setup fiber
	app := fiber.New(utils.NewFiberError())
	app.Use(recover.New())
	app.Use(cors.New())

	// Register the request ID and logging middleware
	//app.Use(middlewares.AuthMiddleware()) // if you want all handler go through
	app.Use(middlewares.RequestID())
	app.Use(middlewares.LogRequestResponse())

	// Initialize the database connection
	if err := database.Connect(config.AppConfig); err != nil {
		logruspack.Logger.Fatalf("Failed to connect to the database: %v", err)
	}

	// Setup routes
	v1.SetupAuthRoutes(app, database.DB, config.AppConfig)
	v1.SetupUserRoutes(app, database.DB, config.AppConfig)

	//App start listen port
	err := app.Listen(":" + strconv.Itoa(config.AppConfig.Server.Port))
	if err != nil {
		logruspack.Logger.Fatalf("Failed to connect to the port : %v", err)
	}
}
