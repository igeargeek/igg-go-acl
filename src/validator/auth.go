package validator

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/igeargeek/igg-golang-api-response/response"
)

type Login struct {
	Username string `validate:"email"`
	Password string
}

func LoginValidator(c *fiber.Ctx) error {
	body := new(Login)
	if err := c.BodyParser(body); err != nil {
		status, resData := response.InternalServerError("")
		return c.Status(status).JSON(resData)
	}
	validate := validator.New()
	err := validate.Struct(body)
	errs := validationErrorFormat(err)
	if len(errs) > 0 {
		status, resData := response.ValidateFailed([]interface{}{errs}, "")
		return c.Status(status).JSON(resData)
	}
	return c.Next()
}
