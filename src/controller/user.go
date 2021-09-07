package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igeargeek/igg-golang-api-response/response"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/igeargeek/igg-go-acl/model"
	"github.com/igeargeek/igg-go-acl/validator"
)

type userController interface {
	GetUsers(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUserByID(c *fiber.Ctx) error
}

type UserController struct{}

func NewUser() userController {
	return &UserController{}
}

func (u *UserController) GetUsers(c *fiber.Ctx) error {
	type Query struct {
		Limit int64 `query:"limit"`
		Page  int64 `query:"page"`
	}
	query := new(Query)
	if err := c.QueryParser(query); err != nil {
		return err
	}
	userModel := model.NewUser()
	results, err := userModel.GetUserPaginate(query.Limit, query.Page, bson.D{}, "-updateAt")
	if err != nil {
		status, resData := response.InternalServerError("")
		return c.Status(status).JSON(resData)
	}
	status, resData := response.Paginate(results, "")
	return c.Status(status).JSON(resData)
}

func (u *UserController) GetUserById(c *fiber.Ctx) error {
	userModel := model.NewUser()
	result, err := userModel.GetUserById(c.Params("id"))
	if err != nil {
		status, resData := response.NotFound("")
		return c.Status(status).JSON(resData)
	}
	status, resData := response.Item(result, "")
	return c.Status(status).JSON(resData)
}

func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var body validator.UserCreate
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	userModel := model.NewUser()
	user := model.User{
		Username: body.Username,
		Password: body.Password,
		Fullname: body.Fullname,
		Roles:    body.Roles,
	}
	err := userModel.CreateUser(&user)
	if err != nil {
		status, resData := response.UnprocessableEntity("")
		return c.Status(status).JSON(resData)
	}
	status, resData := response.Created(user, "")
	return c.Status(status).JSON(resData)
}

func (u *UserController) UpdateUserByID(c *fiber.Ctx) error {
	var body validator.UserUpdate
	if err := c.BodyParser(&body); err != nil {
		status, resData := response.InternalServerError("")
		return c.Status(status).JSON(resData)
	}
	userModel := model.NewUser()
	_, err := userModel.GetUserById(c.Params("id"))
	if err != nil {
		status, resData := response.NotFound("")
		return c.Status(status).JSON(resData)
	}

	upsetData := bson.M{
		"fullname": body.Fullname,
		"roles":    body.Roles,
	}
	if body.Password != "" {
		upsetData["password"] = body.Password
	}
	if err := userModel.UpdateUserByID(upsetData, c.Params("id")); err != nil {
		status, resData := response.UnprocessableEntity("")
		return c.Status(status).JSON(resData)
	}
	status, resData := response.Updated(nil, "")
	return c.Status(status).JSON(resData)
}
