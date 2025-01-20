package ping

import (
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) ping(c *fiber.Ctx) error {

	ping := r.service.Ping()

	return utils.SuccessResponse(c, ping, nil)
}
