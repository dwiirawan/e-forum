package routes

import (
  "libs/api-core/features/vote/dto"
  "libs/api-core/utils"

  "github.com/gofiber/fiber/v2"
)

func (r *Route) Vote(c *fiber.Ctx) error {

  payload := dto.Vote{}

  if err := c.BodyParser(&payload); err != nil {
    return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request body", err)
  }

  user := r.server.Auth.GetUser(c)

  err := r.service.Vote(payload, user.ID)

  if err != nil {

    return err
  }

  return utils.CreatedResponse(c, "created")
}
