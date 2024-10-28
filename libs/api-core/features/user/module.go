package user

import (
	router "libs/api-core/features/user/routes"
	server "libs/api-core/server"
)

type UserModule struct{}

func (*UserModule) Init(server *server.WebServer) {
	router.New(server)
}

func Module() *UserModule {
	return &UserModule{}
}
