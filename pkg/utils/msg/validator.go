package msg

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/mail"
	"strconv"
	"strings"
)

// Validate validates the given model using the custom validator.
func Validate(model interface{}) {
	validate := validator.New()
	validate.RegisterValidation("email", validateEmail)
	validate.RegisterValidation("passwd", validatePassword)

	err := validate.Struct(model)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": generateErrorMessage(err),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		if errJson != nil {
			PanicLogging(errJson)
		}

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

// validatePassword checks if the password meets the required criteria.
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Default min and max values
	min := 8
	max := 50

	// Extract min and max parameters if provided
	params := strings.Split(fl.Param(), ",")
	if len(params) > 0 && params[0] != "" {
		if parsedMin, err := strconv.Atoi(params[0]); err == nil {
			min = parsedMin
		}
	}
	if len(params) > 1 && params[1] != "" {
		if parsedMax, err := strconv.Atoi(params[1]); err == nil {
			max = parsedMax
		}
	}

	var (
		hasMinLen  = len(password) >= min
		hasMaxLen  = len(password) <= max
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*", char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// generateErrorMessage generates a user-friendly error message based on the validation tag.
func generateErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "this field is required"
	case "min":
		return "this field must be at least " + err.Param() + " characters long"
	case "max":
		return "this field must be no more than " + err.Param() + " characters long"
	case "email":
		return "this field must be a valid email address"
	case "passwd":
		return "the password must be between " + err.Param() + " and " + err.Param() + " characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character"
	default:
		return "this field is invalid"
	}
}
