package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Public  fiber.Router
	Private fiber.Router
}

func New(app *fiber.App, authMid fiber.Handler) *Router {
	pathRoot := "/api"
	public := PublicRoute(app, pathRoot)
	private := PrivateRoute(app, pathRoot, authMid)
	router := &Router{
		Public:  public,
		Private: private,
	}

	return router
}

// Public routes don't require authentication
func PublicRoute(app *fiber.App, pathRoot string) fiber.Router {
	route := app.Group(pathRoot)

	return route

	// Add more public routes here
}

// Private routes require authentication
func PrivateRoute(app *fiber.App, pathRoot string, authMid fiber.Handler) fiber.Router {
	route := app.Group(pathRoot, authMid)

	return route
}
