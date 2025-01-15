package router

import (
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) List(c *fiber.Ctx) error {
	parentID := c.Params("parentId")
	res, err := r.service.List(parentID)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(c, res, nil)
}
