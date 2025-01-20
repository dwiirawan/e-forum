package server

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Public  fiber.Router
	Private fiber.Router
}

// Public routes don't require authentication
func PublicRoute(app *fiber.App, prefix string) fiber.Router {
	return app.Group(prefix)
}

// Private routes require authentication
func PrivateRoute(app *fiber.App, prefix string, authMid fiber.Handler) fiber.Router {
	return app.Group(prefix, authMid)
}
