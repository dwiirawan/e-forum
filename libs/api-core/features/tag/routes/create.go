package router

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/tag/dto"
	"libs/api-core/utils"
)

func (r *Route) Create(c *fiber.Ctx) error {
	payload := new(dto.TagCreate)
	if err := c.BodyParser(payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request", err)
	}
	res, err := r.service.Create(*payload)
	if err != nil {
		return err
	}
	return utils.CreatedResponse(c, map[string]interface{}{"id": res})
}
