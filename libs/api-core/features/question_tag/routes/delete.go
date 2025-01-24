package routes

import (
	"libs/api-core/features/question_tag/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Delete(c *fiber.Ctx) error {

	payload := dto.DeleteQuestionTagDto{}

	if err := c.BodyParser(&payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request body", err)
	}

	err := r.service.Delete(payload)

	if err != nil {
		return err
	}

	return utils.DeletedResponse(c, nil)
}
