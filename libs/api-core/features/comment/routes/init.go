package router

import (
	"libs/api-core/features/comment/services"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service *services.CommentService
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PrivateApi("comment"),
		service: services.NewCommentService(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Post("", r.Create)
	r.router.Delete(":id", r.Delete)
	r.router.Get("list/:parentId", r.List)
}
