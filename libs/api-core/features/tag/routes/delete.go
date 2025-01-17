package router

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/utils"
)

func (r *Route) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := r.service.Delete(id)
	if err != nil {
		return err
	}
	return utils.DeletedResponse(c, nil)
}
