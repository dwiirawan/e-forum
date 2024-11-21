package auth

import (
	auth "libs/api-core/features/auth/services"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service *auth.AuthService
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PublicApi("auth"),
		service: auth.NewAuthService(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Post("/register", r.registerUser)
	r.router.Post("/login", r.login)
}
