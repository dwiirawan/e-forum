package router

import (
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := r.service.Delete(id)
	if err != nil {
		return err
	}
	return utils.DeletedResponse(c, nil)
}
