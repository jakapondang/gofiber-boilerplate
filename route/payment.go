package route

import (
	"github.com/gofiber/fiber/v3"
	"goamartha/controller"
	"goamartha/repository/repositoryimpl"
	"goamartha/service/serviceimpl"
	"gorm.io/gorm"
)

func Payment(app *fiber.App, db *gorm.DB) {
	repository := repositoryimpl.NewPaymentRespositoryImpl(db)
	brepository := repositoryimpl.NewBorrowerRepositoryImpl(db)
	service := serviceimpl.NewPaymentServiceImpl(&repository)
	bservice := serviceimpl.NewBorrowerServiceImpl(&brepository)

	controller := controller.NewPaymentController(&bservice, &service)

	app.Post("/payment/:id", controller.MakePayment)
	app.Post("/repayment/:id", controller.MakeRePayment)
}
