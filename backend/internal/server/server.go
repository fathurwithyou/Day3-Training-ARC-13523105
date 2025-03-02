package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/handlers"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)

	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)
}
