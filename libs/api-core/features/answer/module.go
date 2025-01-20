package answer

import (
	router "libs/api-core/features/answer/routes"
	"libs/api-core/server"
)

type AnswerModule struct{}

func (*AnswerModule) Init(server *server.WebServer) {
	router.New(server)
}
