package main

import (
	"fmt"
	"libs/api-core/common"
	"libs/api-core/server"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listEnv := utils.LoadEnv(1)
	apps := server.New(listEnv.APP_NAME)

	client := common.NewBasicJwtAuth(listEnv.JWT_SECRET)

	apps.UseAuth(client, nil)

	apps.RootApiPrefix = "api/v1"

	public := apps.PublicApi("public")
	public.Get("", func(c *fiber.Ctx) error {
		return c.JSON("fooo")
	})

	private := apps.PrivateApi("private")
	private.Get("", func(c *fiber.Ctx) error {
		return c.JSON("fooo")
	})

	port := fmt.Sprintf(":%d", listEnv.APP_PORT)
	apps.App.Listen(port)

}
