package user

import (
	service "libs/api-core/features/user/services"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service service.Service
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PublicApi("user"),
		service: service.New(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Get("/:id", r.getUser)
}
