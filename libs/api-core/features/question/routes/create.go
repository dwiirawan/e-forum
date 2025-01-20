package router

import (
	"libs/api-core/features/question/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Create(c *fiber.Ctx) error {

	payload := dto.QuestionCreate{}

	if err := c.BodyParser(&payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request body", err)
	}

	user := r.server.Auth.GetUser(c)

	err := r.service.Create(payload, user.ID)

	if err != nil {
		return err
	}

	return utils.CreatedResponse(c, "created")
}
