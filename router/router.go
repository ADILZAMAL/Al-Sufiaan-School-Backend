package router

import (
	"al-sufiaan-school-backend/handler"
	// "al-sufiaan-school-backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	//School
	school := api.Group("/school")
	school.Post("/onboard", handler.OnboardSchool)

	// User
	user := api.Group("/user")
	// user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	// user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	// product := api.Group("/product")
	// product.Get("/", handler.GetAllProducts)
	// product.Get("/:id", handler.GetProduct)
	// product.Post("/", middleware.Protected(), handler.CreateProduct)
	// product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
