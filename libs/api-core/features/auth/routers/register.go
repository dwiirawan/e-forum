package auth

import (
	"github.com/gofiber/fiber/v2"
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/utils"
)

func (r *Route) registerUser(c *fiber.Ctx) error {

	payload := new(auth.RegisterUserRequestDTO)

	if err := c.BodyParser(&payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request", err)
	}

	err := r.service.RegisterUser(payload)
	if err != nil {
		return err
	}
	return utils.CreatedResponse(c, nil)
}
