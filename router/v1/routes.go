package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	controller "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/controllers/api/v1"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/middleware"
)

func Routes(app *fiber.App) {
	api := app.Group("/api") // /api

	v1 := api.Group("/v1") // /api/v1
	v1.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	v1.Use(helmet.New())

	v1.Get("/user", middleware.Protected(), controller.AllUsers)
	v1.Post("/user", controller.RegisterUser)
	v1.Post("/auth", controller.LoginUser)

	v1.Post("/contact", controller.CreateContactInfo)
	v1.Get("/contact", controller.GetContactInfo)

	v1.Post("/membership", controller.CreateMemberShip)
	v1.Get("/membership", controller.GetMemberships)
	// v1.Post("/contact")
}
