package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/fathurwithyou/Day3-Training-ARC-13523105/backend/internal/server"
)

func main() {
	app := fiber.New()

	server.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
