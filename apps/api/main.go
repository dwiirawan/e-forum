package main

import (
	"fmt"
	"libs/api-core/middleware"
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

	router := middleware.NewRouter(app.Fiber, middleware.BasicAuthMiddleware())

	router.Public.Get("/public", func(c *fiber.Ctx) error {
		return c.SendString("Public route")
	})

	port := fmt.Sprintf(":%d", listEnv.APP_PORT)
	app.Fiber.Listen(port)

}
