package routes

import (
	"libs/api-core/features/vote/services"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service *services.VoteService
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PrivateApi("vote"),
		service: services.NewVoteService(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Post("", r.Vote)
	r.router.Delete("", r.UnVote)
}
