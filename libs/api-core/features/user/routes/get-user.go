package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserParams struct {
	ID string `params:"id"`
}

func (r *Route) getUser(c *fiber.Ctx) error {
	params := UserParams{}

	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := r.service.GetUser(params.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}
