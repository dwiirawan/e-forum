package router

import (
	"libs/api-core/features/question/services"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service *services.QuestionService
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PrivateApi("question"),
		service: services.NewQuestionService(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Post("", r.Create)
	r.router.Get(":id", r.Get)
}
