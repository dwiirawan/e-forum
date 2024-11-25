package auth

import (
	"github.com/gofiber/fiber/v2"
	dto "libs/api-core/features/auth/dto"
	"libs/api-core/utils"
)

func (r *Route) login(c *fiber.Ctx) error {
	payload := new(dto.LoginRequestDTO)
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request", err)
	}

	res, err := r.service.Login(payload)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, res)
}
