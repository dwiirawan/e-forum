package router

import (
	"fmt"
	"libs/api-core/features/answer/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Print(id, "<<< iki ID ne")
	payload := new(dto.AnswerUpdate)
	if err := c.BodyParser(payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_INVALID_PAYLOAD", "invalid payload", err)
	}

	err := r.service.Update(id, *payload)
	if err != nil {
		return err
	}
	return utils.UpdatedResponse(c, map[string]string{"id": id})
}
