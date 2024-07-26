package users

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Route(app *fiber.App, db *gorm.DB) {
	repository := NewRepositoryImpl(db)
	service := NewServiceImpl(&repository)
	controller := NewController(&service)
	app.Get("/users", controller.FindAll)
	app.Post("/users", controller.Create)

}
