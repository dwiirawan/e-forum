package router

import (
	"errors"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Get(c *fiber.Ctx) error {

	id := c.Params("id")

	if len(id) == 0 {
		return utils.NewError(fiber.StatusBadRequest, "PARSER GET ID QUESTION DETAIL", "id is not fill", errors.New("id is not fill"))
	}

	res, err := r.service.Get(id)

	if err != nil {
		return err
	}

	return utils.SuccessResponse(c, res, nil)

}
