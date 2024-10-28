package server

import (
	"fmt"
	"libs/api-core/middleware"

	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	App           *fiber.App
	Auth          *middleware.WebAuthManager
	RootApiPrefix string
}

func New(appName string, auth *middleware.WebAuthManager) *WebServer {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       appName,
	})

	return &WebServer{
		App:  app,
		Auth: auth,
	}
}

func (s WebServer) PublicApi(prefix string) fiber.Router {
	fullPrefix := fmt.Sprintf("/%s/%s", s.RootApiPrefix, prefix)
	return PublicRoute(s.App, fullPrefix)
}

func (s WebServer) PrivateApi(prefix string) fiber.Router {
	fullPrefix := fmt.Sprintf("/%s/%s", s.RootApiPrefix, prefix)
	return s.App.Group(fullPrefix, s.Auth.AuthGuardMiddleware)
}
