package v1c

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/database"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"
	v1r "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/resources/api/v1"
	v1service "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/services/api/v1"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/utils"
)

func AllUsers(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["fullname"].(string)
	var users []models.User
	database.Database.DB.Find(&users)

	return c.Status(200).JSON(map[string]interface{}{"name": name, "users": users})
}
func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: err}))
	}
	errors := utils.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: errors}))
	}
	hashPassword, _ := utils.HashPassword((user.Password))
	user.Password = hashPassword

	response := v1service.CreateUser(user)
	if _, errormessage := response["error"]; errormessage {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: response["error"]}))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.Response(utils.ResponseData{Data: response["data"], Message: "Account has created"}))
}

func LoginUser(c *fiber.Ctx) error {
	var input v1r.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: err}))
	}
	errors := utils.ValidateStruct(input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: errors}))
	}
	response := v1service.AccessToken(input)
	if _, errormessage := response["error"]; errormessage {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response(utils.ResponseData{Error: response["error"]}))
	}
	return c.Status(fiber.StatusAccepted).JSON(utils.Response(utils.ResponseData{Data: fiber.Map{"token": response["token"], "user": response["user"]}, Message: "Your access token"}))
}
