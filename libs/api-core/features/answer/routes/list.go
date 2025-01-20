package router

import (
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) List(c *fiber.Ctx) error {
	questionID := c.Params("questionId")
	res, err := r.service.List(questionID)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(c, res, nil)
}
