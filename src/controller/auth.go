package controller

import (
	"os"
	"strings"
	"time"

	"github.com/igeargeek/igg-go-acl/model"
	"github.com/igeargeek/igg-go-acl/validator"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/igeargeek/igg-golang-api-response/response"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type authController interface {
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type Auth struct{}

func NewAuth() authController {
	return &Auth{}
}

func generateToken(u model.User) (string, error) {
	u.Password = ""
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.Id
	claims["roles"] = []string{u.Roles}
	claims["username"] = u.Username
	claims["fullname"] = u.Fullname
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return t, err
}

func (u *Auth) Login(c *fiber.Ctx) error {
	var body validator.Login
	if err := c.BodyParser(&body); err != nil {
		status, resData := response.InternalServerError("")
		return c.Status(status).JSON(resData)
	}
	userModel := model.NewUser()
	user, _ := userModel.GetUserOne(bson.M{
		"username": body.Username,
	})
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		status, resData := response.Unauthorized("")
		return c.Status(status).JSON(resData)
	}

	t, err := generateToken(user)
	if err != nil {
		status, resData := response.InternalServerError("")
		return c.Status(status).JSON(resData)
	}
	status, resData := response.Success(bson.M{"accessToken": t}, "")
	return c.Status(status).JSON(resData)
}

func (u *Auth) RefreshToken(c *fiber.Ctx) error {
	tokenString := strings.Replace(c.Get("Authorization"), "Bearer ", "", 1)
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		status, resData := response.Unauthorized("")
		return c.Status(status).JSON(resData)
	}

	claims := token.Claims.(jwt.MapClaims)
	data := claims["user"].(map[string]interface{})
	username := data["Username"].(string)

	userModel := model.NewUser()
	user, _ := userModel.GetUserOne(bson.M{"username": username})
	t, err := generateToken(user)
	if err != nil {
		status, resData := response.InternalServerError("")
		return c.Status(status).JSON(resData)
	}
	status, resData := response.Success(bson.M{"accessToken": t}, "")
	return c.Status(status).JSON(resData)
}
