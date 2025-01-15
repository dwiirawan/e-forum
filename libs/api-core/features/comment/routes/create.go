package router

import (
	"libs/api-core/features/comment/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) Create(c *fiber.Ctx) error {

	payload := dto.CommentCreate{}

	if err := c.BodyParser(&payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request body", err)
	}

	user := r.server.Auth.GetUser(c)

	commentID, err := r.service.Create(payload, user.ID)

	if err != nil {
		return err
	}

	return utils.CreatedResponse(c, map[string]string{"id": *commentID})
}
