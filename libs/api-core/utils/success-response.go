package utils

import "github.com/gofiber/fiber/v2"

type SuccessDto struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Data    *interface{} `json:"data"`
	Meta    interface{}  `json:"meta"`
}

func newSuccess(ctx *fiber.Ctx, statusCode int, msg string, data *interface{}) error {
	return ctx.Status(statusCode).JSON(&SuccessDto{
		Code:    statusCode,
		Message: msg,
		Data:    data,
	})
}

func CreatedResponse(ctx *fiber.Ctx, data *interface{}) error {
	return newSuccess(ctx, fiber.StatusCreated, "Created", data)
}

func UpdatedResponse(ctx *fiber.Ctx, data *interface{}) error {
	return newSuccess(ctx, fiber.StatusOK, "Updated", data)
}
