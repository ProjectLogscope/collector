package middleware

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/hardeepnarang10/collector/service/handler/v1/message"
)

func (m *middleware) Collect(c *fiber.Ctx) error {
	var logMessage message.Log
	if err := c.BodyParser(&logMessage); err != nil {
		slog.DebugContext(c.Context(), "unable to parse incoming request body to log message type", slog.Any("error", err))
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	if errs := m.validator.Validate(logMessage); len(errs) > 0 && errs[0].Error {
		errMap := make(map[string]string, 0)
		for _, err := range errs {
			errMap[err.FailedField] = fmt.Sprintf("'%+v'; constraint: %q", err.Value, err.Tag)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errMap)
	}

	return c.Next()
}
