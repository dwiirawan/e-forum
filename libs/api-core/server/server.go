package server

import (
	"fmt"
	"libs/api-core/middleware"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App           *fiber.App
	Auth          middleware.WebAuthManager
	RootApiPrefix string
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
		App: app,
	}
}

func (s Server) PublicApi(prefix string) fiber.Router {
	fullPrefix := fmt.Sprintf("/%s/%s", s.RootApiPrefix, prefix)
	return PublicRoute(s.App, fullPrefix)
}

func (s Server) PrivateApi(prefix string) fiber.Router {
	fullPrefix := fmt.Sprintf("/%s/%s", s.RootApiPrefix, prefix)
	return PrivateRoute(s.App, fullPrefix, s.Auth.AuthGuardMiddleware)
}

func (s Server) UseAuth(client middleware.WebAuthClient, bearerTokenConfig *middleware.BearerTokenMiddlewareConfig) {
	s.Auth = *middleware.NewWebAuthManager(client, bearerTokenConfig)
}
