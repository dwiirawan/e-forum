package auth

import (
	"github.com/gofiber/fiber/v2"
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/utils"
)

func (r *Route) registerUser(c *fiber.Ctx) error {

	payload := new(auth.RegisterUserRequestDTO)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := r.service.RegisterUser(payload)
	if err != nil {
		return err
	}
	return utils.CreatedResponse(c, nil)
}
