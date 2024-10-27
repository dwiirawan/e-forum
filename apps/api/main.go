package main

import (
	"fmt"
	"libs/api-core/server"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listEnv := utils.LoadEnv(1)
	app := server.New(listEnv.APP_NAME)

	app.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := fmt.Sprintf(":%d", listEnv.APP_PORT)
	app.Fiber.Listen(port)

}
