package question_tag

import (
	router "libs/api-core/features/question_tag/routes"
	"libs/api-core/server"
)

type QuestionTagModule struct{}

func (*QuestionTagModule) Init(server *server.WebServer) {
	router.New(server)
}
