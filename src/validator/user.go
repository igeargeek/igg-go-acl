package validator

import (
	"github.com/igeargeek/igg-go-acl/model"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/igeargeek/igg-golang-api-response/response"
	"go.mongodb.org/mongo-driver/bson"
)

type UserCreate struct {
	Username string `validate:"email,max=50"`
	Password string `validate:"min=6,max=20"`
	Fullname string `validate:"required"`
	Roles    string `validate:"required"`
}

type UserUpdate struct {
	Password string
	Fullname string `validate:"required"`
	Roles    string `validate:"required"`
}

func CreateUserValidator(c *fiber.Ctx) error {
	body := new(UserCreate)
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

	userModel := model.NewUser()
	result, err := userModel.GetUserOne(bson.M{"username": body.Username})
	if result.Username != "" {
		status, resData := response.BadRequest("Username exist")
		return c.Status(status).JSON(resData)
	}
	return c.Next()
}

func UpdateUserValidator(c *fiber.Ctx) error {
	body := new(UserUpdate)
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
