package definition

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary 			Insert log message
// @Description 	Insert a new log record
// @Tags 					Collect
// @Accept 				json
// @Param 				filter	body	Log	true	"Log collect parameters"
// @Produce 			text/plain
// @Success 			200			"OK"
// @Router 				/				[post]
func Collect(c *fiber.Ctx) error
