package router

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/utils"
)

func (r *Route) List(c *fiber.Ctx) error {
	res, err := r.service.List()
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, res, nil)
}
