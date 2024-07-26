package route

import (
	"github.com/gofiber/fiber/v3"
	"goamartha/controller"
	"goamartha/repository/repositoryimpl"
	"goamartha/service/serviceimpl"
	"gorm.io/gorm"
)

func Borrower(app *fiber.App, db *gorm.DB) {
	repository := repositoryimpl.NewBorrowerRepositoryImpl(db)
	brepository := repositoryimpl.NewPaymentRespositoryImpl(db)
	service := serviceimpl.NewBorrowerServiceImpl(&repository)
	pservice := serviceimpl.NewPaymentServiceImpl(&brepository)
	controller := controller.NewBorrowerController(&service, &pservice)
	app.Get("/borrow", controller.FindAll)
	app.Get("/borrow/:id", controller.FindById)
	app.Post("/borrow", controller.Create)
	app.Put("/borrow/:id", controller.UpdateTrxDate)
}
