package router

import (
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/igeargeek/igg-go-acl/controller"
	"github.com/igeargeek/igg-go-acl/validator"
	fibercasbinrest "github.com/prongbang/fiber-casbinrest"
)

func V1(app *fiber.App) {

	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	app.Use(fibercasbinrest.NewDefault(e, os.Getenv("SECRET_KEY")))

	v1 := app.Group("/v1")

	userController := controller.NewUser()
	users := v1.Group("/users")
	users.Get("/", userController.GetUsers)
	users.Get("/:id", userController.GetUserById)
	users.Post("/", validator.CreateUserValidator, userController.CreateUser)
	users.Put("/:id", validator.UpdateUserValidator, userController.UpdateUserByID)

	authController := controller.NewAuth()
	auth := v1.Group("/auth")
	auth.Post("/login", validator.LoginValidator, authController.Login)
	auth.Get("/refresh-token", authController.RefreshToken)
}
