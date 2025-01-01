package router

import (
	"libs/api-core/features/question/dto"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (r *Route) List(c *fiber.Ctx) error {
	payload := new(dto.PaginationListQuestionRequest)
	if err := c.QueryParser(payload); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "E_BAD_REQUEST", "invalid request", err)
	}
	res, meta, err := r.service.List(payload)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(c, res, meta)
}
