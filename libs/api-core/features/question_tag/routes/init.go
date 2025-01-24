package routes

import (
	"libs/api-core/features/question_tag/services"
	"libs/api-core/server"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	server  *server.WebServer
	router  fiber.Router
	service *services.QuestionTagService
}

func New(server *server.WebServer) *Route {
	route := Route{
		server:  server,
		router:  server.PrivateApi("question-tag"),
		service: services.NewQuestionTagService(server.DB),
	}
	route.register()
	return &route
}

func (r *Route) register() {
	r.router.Post("", r.Create)
	r.router.Delete(":id", r.Delete)
}
