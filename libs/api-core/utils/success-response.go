package utils

import "github.com/gofiber/fiber/v2"

type SuccessDto struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Meta    any    `json:"meta"`
}

func newSuccess(ctx *fiber.Ctx, statusCode int, msg string, data interface{}, meta any) error {
	return ctx.Status(statusCode).JSON(&SuccessDto{
		Code:    statusCode,
		Message: msg,
		Data:    data,
		Meta:    meta,
	})
}

func CreatedResponse(ctx *fiber.Ctx, data interface{}) error {
	return newSuccess(ctx, fiber.StatusCreated, "Created", data, nil)
}

func UpdatedResponse(ctx *fiber.Ctx, data *any) error {
	return newSuccess(ctx, fiber.StatusOK, "Updated", data, nil)
}

func DeletedResponse(ctx *fiber.Ctx, data *any) error {
	return newSuccess(ctx, fiber.StatusOK, "Deleted", data, nil)
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}, meta any) error {
	return newSuccess(ctx, fiber.StatusOK, "OK", data, meta)
}
