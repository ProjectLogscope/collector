package middleware

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	Collect(*fiber.Ctx) error
}
