package vote

import (
	router "libs/api-core/features/vote/routes"
	"libs/api-core/server"
)

type VoteModule struct{}

func (*VoteModule) Init(server *server.WebServer) {
	router.New(server)
}
