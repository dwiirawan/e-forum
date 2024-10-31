package server

import (
	"fmt"
	"gorm.io/gorm"
	"libs/api-core/database"
	"libs/api-core/middleware"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	App           *fiber.App
	Auth          *middleware.WebAuthManager
	DB            *gorm.DB
	RootApiPrefix string
}

func New(appName string, auth *middleware.WebAuthManager, env utils.Env) *WebServer {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       appName,
	})

	db := database.ConnectDatabasePostgres(database.DBConfigGorm{
		Port:     env.DB_PORT,
		Host:     env.DB_HOST,
		Password: env.DB_PASSWORD,
		User:     env.DB_USER,
		DBName:   env.DB_NAME,
	})

	return &WebServer{
		App:  app,
		Auth: auth,
		DB:   db,
	}
}

func (s WebServer) PublicApi(prefix string) fiber.Router {
	fullPrefix := fmt.Sprintf("/%s/%s", s.RootApiPrefix, prefix)
	return PublicRoute(s.App, fullPrefix)
}

func (s WebServer) PrivateApi(prefix string) fiber.Router {
	fullPrefix := fmt.Sprintf("/%s/%s", s.RootApiPrefix, prefix)
	return s.App.Group(fullPrefix, s.Auth.AuthGuardMiddleware)
}
