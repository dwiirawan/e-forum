package user

import (
	"fmt"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server *server.WebServer
	router fiber.Router
}

func New(server *server.WebServer) *Route {
	route := Route{
		server: server,
		router: server.PrivateApi("user"),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Get("/", func(c *fiber.Ctx) error {
		fo := r.server.Auth.User(c)
		fmt.Println(fo)
		return nil
	})
}
