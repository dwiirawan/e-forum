package routes

import (
	"libs/api-core/features/vote/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) UnVote(c *fiber.Ctx) error {

	payload := dto.UnVote{}

	if err := c.BodyParser(&payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request body", err)
	}

	user := r.server.Auth.GetUser(c)

	err := r.service.UnVote(payload, user.ID)

	if err != nil {
		return err
	}

	return utils.DeletedResponse(c, nil)
}
