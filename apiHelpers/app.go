package apiHelpers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	route "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/router/v1"
)

func Init() {
	app := fiber.New(fiber.Config{})
	app.Use(recover.New())
	file, err := os.OpenFile("./logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))
	route.Routes(app)
	app.Listen(":3001")
}
