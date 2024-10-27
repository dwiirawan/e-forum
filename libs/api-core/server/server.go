package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Fiber *fiber.App
}

func New(appName string) *Server {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       appName,
	})

	return &Server{
		Fiber: app,
	}
}
