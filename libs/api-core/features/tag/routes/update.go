package router

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/tag/dto"
	"libs/api-core/utils"
)

func (r *Route) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	payload := new(dto.TagUpdate)
	if err := c.BodyParser(payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request", err)
	}
	err := r.service.Update(id, *payload)
	if err != nil {
		return err
	}
	return utils.UpdatedResponse(c, "updated")
}
