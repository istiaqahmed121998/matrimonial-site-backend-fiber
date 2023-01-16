package v1c

import (
	"github.com/gofiber/fiber/v2"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/models"
	v1s "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/services/api/v1"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/utils"
)

func CreateMemberShip(c *fiber.Ctx) error {
	membership := new(models.Membership)
	if err := c.BodyParser(membership); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: err}))
	}
	if errors := utils.ValidateStruct(*membership); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(utils.ResponseData{Error: errors}))
	}
	response := v1s.MembershipService{Membership: *membership}.CreateMemberShip()
	if _, errormessage := response["error"]; errormessage {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response(utils.ResponseData{Error: response["error"]}))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.Response(utils.ResponseData{Data: response["data"]}))
}
func GetMemberships(c *fiber.Ctx) error {
	response := v1s.MembershipService{}.GetMemberships()
	if _, errormessage := response["error"]; errormessage {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Response(utils.ResponseData{Error: response["error"]}))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.Response(utils.ResponseData{Data: response["data"]}))
}
