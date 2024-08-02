package msg

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/mail"
)

func Validate(model interface{}) {

	validate := validator.New()
	validate.RegisterValidation("email", validateEmail)
	err := validate.Struct(model)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		PanicLogging(errJson)

		panic(ValidationError{
			Message: string(jsonMessage),
		})
	}
}

// validateEmail checks if the email is valid.
func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	_, err := mail.ParseAddress(email)
	return err == nil
}
