package tag

import (
	router "libs/api-core/features/tag/routes"
	"libs/api-core/server"
)

type TagModule struct{}

func (*TagModule) Init(server *server.WebServer) {
	router.New(server)
}
