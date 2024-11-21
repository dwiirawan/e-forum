package main

import (
	"fmt"
	"libs/api-core/common"
	"libs/api-core/features/auth"
	"libs/api-core/features/user"
	"libs/api-core/middleware"
	"libs/api-core/server"
	"libs/api-core/utils"
)

func main() {
	listEnv := utils.LoadEnv(1)

	client := common.NewBasicJwtAuth()

	webAuth := middleware.NewWebAuthManager(client, nil)

	// init server
	apps := server.New(listEnv.APP_NAME, webAuth, listEnv)

	apps.RootApiPrefix = "api/v1"

	apps.UseModules(&user.UserModule{}, &auth.AuthModule{})

	// listen server
	port := fmt.Sprintf(":%d", listEnv.APP_PORT)
	err := apps.App.Listen(port)
	if err != nil {
		panic(err)
	}

}
