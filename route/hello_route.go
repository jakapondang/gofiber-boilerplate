package route

import (
	"github.com/gofiber/fiber/v3"
	"goamartha/controller"
	"goamartha/repository/repositoryimpl"
	"goamartha/service/serviceimpl"
	"gorm.io/gorm"
)

func Hello(app *fiber.App, db *gorm.DB) {
	helloRepository := repositoryimpl.NewHelloRepositoryImpl(db)
	helloService := serviceimpl.NewHelloServiceImpl(&helloRepository)
	helloController := controller.NewHelloController(&helloService)
	app.Get("/:id", helloController.Pageworld)
}
