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
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
	"gofiber-boilerplatev3/pkg/utils/logruspack"
	"strconv"
)

func main() {
	// Initialize Logrus
	logruspack.Init()

	// Load configurations
	config.Init()

	// Initialize JWT settings
	jwt.SetConfig(
		config.AppConfig.JWT.Secret,
		config.AppConfig.JWT.AppName,
		config.AppConfig.JWT.Audience,
		config.AppConfig.JWT.ExpAccessToken,
		config.AppConfig.JWT.ExpRefreshToken,
	)
	//setup fiber
	app := fiber.New(utils.NewFiberError())
	app.Use(recover.New())
	app.Use(cors.New())

	// Register the request ID and logging middleware
	//app.Use(middlewares.AuthMiddleware())
	app.Use(middlewares.RequestID())
	app.Use(middlewares.LogRequestResponse())

	// Initialize the database connection
	if err := database.Connect(config.AppConfig); err != nil {
		logruspack.Logger.Fatalf("Failed to connect to the database: %v", err)
	}

	// Setup routes
	v1.SetupUserRoutes(app, database.DB, config.AppConfig)

	//App start listen port
	err := app.Listen(":" + strconv.Itoa(config.AppConfig.Server.Port))
	if err != nil {
		logruspack.Logger.Fatalf("Failed to connect to the port : %v", err)
	}
}
