package common

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type BasicJwtAuth struct {
	SecretKey string
}

func NewBasicJwtAuth(secretKey string) BasicJwtAuth {
	return BasicJwtAuth{
		SecretKey: secretKey,
	}
}

func (b BasicJwtAuth) GetUserFromToken(token string) (any, error) {
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
