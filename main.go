package main

import (
	"al-sufiaan-school-backend/database"
	"al-sufiaan-school-backend/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "al-sufiaan-school-backend/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "*",
	}))

	database.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(":3500")
}
