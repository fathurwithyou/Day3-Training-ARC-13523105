package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/handlers"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	app.Use(middleware.LoggingMiddleware)

	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	app.Get("/examscores/:id", handlers.GetExamScoresByUserID)
	app.Post("/examscores", handlers.CreateExamScore)
	app.Put("/examscores/:id", handlers.UpdateExamScore)
	app.Delete("/examscores/:id", handlers.DeleteExamScore)

	app.Get("/studentcourses", handlers.GetStudentCourses)
}
