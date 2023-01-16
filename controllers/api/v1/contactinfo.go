package v1c

import (
	"github.com/gofiber/fiber/v2"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"
	v1s "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/services/api/v1"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/utils"
)

func CreateContactInfo(c *fiber.Ctx) error {
	contactinfo := new(models.ContactInfo)
	if err := c.BodyParser(contactinfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: err}))
	}
	response := v1s.ContactInfoService{ContactInfo: contactinfo}.CreateContactInfo()
	if _, errormessage := response["error"]; errormessage {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response(utils.ResponseData{Error: response["error"]}))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.Response(utils.ResponseData{Data: response["data"]}))
}

func GetContactInfo(c *fiber.Ctx) error {
	response := v1s.ContactInfoService{}.GetContactInfo()
	if _, errormessage := response["error"]; errormessage {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response(utils.ResponseData{Error: response["error"]}))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.Response(utils.ResponseData{Data: response["data"]}))
}
