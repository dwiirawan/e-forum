package ping

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/utils"
)

func (r *Route) ping(c *fiber.Ctx) error {

	ping := r.service.Ping()

	return utils.SuccessResponse(c, ping)
}
