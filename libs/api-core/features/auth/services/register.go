package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/gofiber/fiber/v2"
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
	"os"
)

func (s *AuthService) RegisterUser(payload *auth.RegisterUserRequestDTO) error {
	hashedPassword, salt, err := s.CreateHashAndSalt(payload.Password)
	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_CREATE_HASH_AND_SALT", "failed to create hash and salt", err)
	}

	userModel := models.UserModel{
		Email:    payload.Email,
		Username: payload.Username,
		Hash:     hashedPassword,
		Salt:     salt,
		Name:     payload.Username,
	}
	err = s.db.Model(userModel).Create(&userModel).Error

	if err != nil {

		errDuplicateUsername := errors.New("ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)")
		if err.Error() == errDuplicateUsername.Error() {
			return utils.NewError(fiber.StatusForbidden, "E_DUPLICATED_KEY_USERNAME", "username already exists", errDuplicateUsername)
		}

		errDuplicateEmail := errors.New("ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")
		if err.Error() == errDuplicateEmail.Error() {
			return utils.NewError(fiber.StatusForbidden, "E_DUPLICATED_KEY_EMAIL", "email already exists", errDuplicateEmail)

		}
		return utils.NewError(fiber.StatusInternalServerError, "E_CREATE_USER", "failed to create user", errDuplicateUsername)
	}

	return nil
}

func (s *AuthService) CreateHashAndSalt(password string) (hashResponse string, saltResponse string, err error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", "", utils.NewError(fiber.StatusInternalServerError, "E_SECRET_KEY", utils.ERR_INTERNAL_SERVER_ERROR, errors.New("SECRET_KEY environment variable not set"))
	}

	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return "", "", utils.NewError(fiber.StatusInternalServerError, "E_CREATE_SALT", utils.ERR_INTERNAL_SERVER_ERROR, errors.New("SALT generation failed"))
	}

	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write(salt)
	hash.Write([]byte(secretKey))

	hashedPassword := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	saltString := base64.StdEncoding.EncodeToString(salt)

	return hashedPassword, saltString, nil
}
