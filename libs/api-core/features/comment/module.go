package comment

import (
	router "libs/api-core/features/comment/routes"
	"libs/api-core/server"
)

type CommentModule struct{}

func (*CommentModule) Init(server *server.WebServer) {
	router.New(server)
}
