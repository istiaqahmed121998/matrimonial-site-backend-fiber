package main

import (
	app "github.com/istiaqahmed121998/matrimonial-site-backend-fiber/apiHelpers"
	"github.com/istiaqahmed121998/matrimonial-site-backend-fiber/database"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.ConnectDb()
	app.Init()
}
