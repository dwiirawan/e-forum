package ping

import (
	"github.com/gofiber/fiber/v2"
	ping "libs/api-core/features/ping/services"
	"libs/api-core/server"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service *ping.PingService
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PrivateApi("ping"),
		service: ping.NewPingService(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Get("/ping", r.ping)

}
