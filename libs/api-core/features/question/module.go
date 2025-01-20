package question

import (
	router "libs/api-core/features/question/routes"
	"libs/api-core/server"
)

type QuestionModule struct{}

func (*QuestionModule) Init(server *server.WebServer) {
	router.New(server)
}
