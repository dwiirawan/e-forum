package common

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/gofiber/fiber/v2"
	"libs/api-core/utils"
	"os"
)

type HashingMethod struct {
	SecretKey string
}

func NewHashingMethod() *HashingMethod {
	return &HashingMethod{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}

func (c *HashingMethod) VerifyPassword(password string, passwordHash string, passwordSalt string) bool {

	hashBytes, err := base64.StdEncoding.DecodeString(passwordHash)
	if err != nil {
		return false
	}
	saltBytes, err := base64.StdEncoding.DecodeString(passwordSalt)
	if err != nil {
		return false
	}
	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write(saltBytes)
	hash.Write([]byte(c.SecretKey))
	hashedPassword := hash.Sum(nil)
	return bytes.Equal(hashBytes, hashedPassword)
}

func (c *HashingMethod) CreateHashAndSalt(password string) (hashResponse string, saltResponse string, err error) {

	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return "", "", utils.NewError(fiber.StatusInternalServerError, "E_CREATE_SALT", utils.ERR_INTERNAL_SERVER_ERROR, errors.New("SALT generation failed"))
	}

	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write(salt)
	hash.Write([]byte(c.SecretKey))

	hashedPassword := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	saltString := base64.StdEncoding.EncodeToString(salt)

	return hashedPassword, saltString, nil
}
