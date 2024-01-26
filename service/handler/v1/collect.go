package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) Collect(c *fiber.Ctx) error {
	if err := h.service.QueueClient.Publish(c.Context(), h.service.ServiceConfig.QueueTopicLog, c.Body()); err != nil {
		slog.ErrorContext(c.Context(),
			"unable to publish incoming request body to queue topic",
			slog.String("body", string(c.Body())),
			slog.String("topic", h.service.ServiceConfig.QueueTopicLog),
			slog.Any("error", err),
		)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
