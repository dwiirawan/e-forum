package middleware

import (
	"libs/api-core/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const USER_LOCAL_KEY = "x-user"

type BearerTokenMiddlewareConfig struct {

	// Optional. Default: "access_token".
	BodyKey string

	// HeaderKey defines the prefix of the Authorization header's value, used when
	// searching for the bearer token inside the request's headers.
	// Optional. Default: "Bearer".
	HeaderKey string

	// QueryKey defines the key to use when searching for the bearer token inside the
	// request's query parameters.
	// Optional. Default: "access_token".
	QueryKey string

	// RequestKey defines the name of the local variable that will be created in the
	// request's context, which will contain the bearer token extracted from the
	// request.
	// Optional. Default: "token".
	RequestKey string
}

type WebAuthClient interface {
	GetUserFromToken(token string) (any, error)
}

type WebAuthManager struct {
	bearerTokenConfig *BearerTokenMiddlewareConfig
	client            WebAuthClient
}

func (m WebAuthManager) AuthGuardMiddleware(ctx *fiber.Ctx) error {
	var token *string

	// get bearer token from request authorization header
	headerValue := ctx.Get("authorization")

	if len(headerValue) > 0 {
		components := strings.SplitN(headerValue, " ", 2)

		if len(components) == 2 && components[0] == m.bearerTokenConfig.HeaderKey {
			token = &components[1]
		}
	} else {

		var queryValue *string
		if m.bearerTokenConfig != nil {
			qValue := ctx.Query(m.bearerTokenConfig.QueryKey)
			queryValue = &qValue
		}

		if queryValue != nil {
			token = queryValue
		}
	}

	if token == nil {
		return ctx.SendStatus(401)
	}

	user, err := m.client.GetUserFromToken(*token)
	if err != nil {
		if err == fiber.ErrUnauthorized {
			return err
		}
		return utils.NewError("WEB_AUTH_MANAGER__GetUserFromToken", err)
	}

	ctx.Locals(USER_LOCAL_KEY, user)
	if user == nil {
		return ctx.SendStatus(401)
	}

	return ctx.Next()

}

func NewWebAuthManager(client WebAuthClient, opts *BearerTokenMiddlewareConfig) *WebAuthManager {
	config := &BearerTokenMiddlewareConfig{
		BodyKey:    "access_token",
		HeaderKey:  "Bearer",
		QueryKey:   "access_token",
		RequestKey: "token",
	}

	if opts != nil {
		if len(opts.BodyKey) > 0 {
			config.BodyKey = opts.BodyKey
		}

		if len(opts.HeaderKey) > 0 {
			config.HeaderKey = opts.HeaderKey
		}

		if len(opts.QueryKey) > 0 {
			config.QueryKey = opts.QueryKey
		}

		if len(opts.RequestKey) > 0 {
			config.RequestKey = opts.RequestKey
		}
	}

	return &WebAuthManager{bearerTokenConfig: config, client: client}
}
