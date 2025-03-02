package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	log.Printf("%s %s %d %v", c.Method(), c.OriginalURL(), c.Response().StatusCode(), duration)
	return err
}
