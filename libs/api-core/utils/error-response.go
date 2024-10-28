package utils

import (
	"fmt"
	"strings"
)

type ErrorDto struct {
	part    string
	err     error
	message string
	// stack   []Trace
}

func NewError(part string, err error) *ErrorDto {
	msg := err.Error()
	parts := []string{}

	if me, ok := err.(*ErrorDto); ok {
		parts = append(parts, me.part)
		msg = me.message
	}

	parts = append(parts, part)

	return &ErrorDto{
		part:    strings.Join(parts, "\r\n"),
		err:     err,
		message: msg,
	}
}

func (e *ErrorDto) Error() string {
	return fmt.Sprintf("%s\r\n%s", e.part, e.message)
}
