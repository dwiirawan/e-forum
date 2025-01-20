package ping

import (
	router "libs/api-core/features/ping/routes"
	"libs/api-core/server"
)

type PingModule struct{}

func (*PingModule) Init(server *server.WebServer) {
	router.New(server)
}
