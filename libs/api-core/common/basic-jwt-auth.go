package common

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"libs/api-core/utils"
	"os"

	"github.com/golang-jwt/jwt"
)

type BasicJwtAuth struct {
	SecretKey string
}

func NewBasicJwtAuth() *BasicJwtAuth {
	secretKey := os.Getenv("JWT_SECRET")
	return &BasicJwtAuth{
		SecretKey: secretKey,
	}
}

func (b *BasicJwtAuth) GetUserFromToken(token string) (any, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(b.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	user := claims["user"]
	if user == nil {
		return nil, fmt.Errorf("user not found in token")
	}

	return user, nil

}

func (b *BasicJwtAuth) GenerateToken(user any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})
	signedToken, err := token.SignedString([]byte(b.SecretKey))
	if err != nil {
		return "", utils.NewError(fiber.StatusInternalServerError, "E_GENERATE_TOKEN", utils.ERR_INTERNAL_SERVER_ERROR, err)
	}
	return signedToken, nil
}
