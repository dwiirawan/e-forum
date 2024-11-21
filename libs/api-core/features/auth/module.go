package auth

import (
	router "libs/api-core/features/auth/routers"
	"libs/api-core/server"
)

type AuthModule struct{}

func (*AuthModule) Init(server *server.WebServer) {
	router.New(server)
}
