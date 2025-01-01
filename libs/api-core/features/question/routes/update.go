package router

import (
	"libs/api-core/features/question/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	payload := new(dto.QuestionUpdate)
	if err := c.BodyParser(payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_INVALID_PAYLOAD", "invalid payload", err)
	}
	err := r.service.Update(id, *payload)
	if err != nil {
		return err
	}
	return utils.UpdatedResponse(c, nil)
}
