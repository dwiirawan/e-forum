package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ErrorDto struct {
	Part    string
	Err     error
	Message string
	ErrCode int
}

const (
	ERR_INTERNAL_SERVER_ERROR = "Internal Server Error"
)

func NewError(errCode int, part string, msg string, err error) *ErrorDto {
	parts := []string{}

	if me, ok := err.(*ErrorDto); ok {
		parts = append(parts, me.Part)
		msg = me.Message
	}

	parts = append(parts, part)

	if errCode == fiber.StatusInternalServerError {
		msg = "Internal Server Error"
	}

	return &ErrorDto{
		Part:    strings.Join(parts, "\r\n"),
		Err:     err,
		Message: msg,
		ErrCode: errCode,
	}
}

func (e *ErrorDto) Error() string {
	return fmt.Sprintf("%s\r\n%s\r\n%s", e.Err.Error(), e.Message, e.Part)
}
