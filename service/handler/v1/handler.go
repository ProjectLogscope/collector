package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Collect(*fiber.Ctx) error
}
