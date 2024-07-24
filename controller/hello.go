package controller

import (
	"github.com/gofiber/fiber/v3"
	"goamartha/domain/model"
	"goamartha/service"
)

type HelloController struct {
	service.HelloService
}

func NewHelloController(helloService *service.HelloService) *HelloController {
	return &HelloController{HelloService: *helloService}
}

//type UserData struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}

func (controller HelloController) Pageworld(c fiber.Ctx) error {
	//data := []UserData{
	//	{"Alice", 30},
	//	{"Bob", 25},
	//}
	//
	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = ioutil.WriteFile("data.json", jsonData, 0644)
	//return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
	//	Code:    200,
	//	Message: "Success",
	//	Data:    "Tell me why",
	//})

	id := c.Params("id")

	result := controller.HelloService.FindById(c.Context(), id)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})

}
