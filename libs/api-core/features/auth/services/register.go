package auth

import (
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/models"
)

func (s *AuthService) RegisterUser(payload *auth.RegisterUserRequestDTO) error {
	hashedPassword, salt, err := s.hashingMethod.CreateHashAndSalt(payload.Password)
	if err != nil {
		return err
	}

	userModel := models.UserModel{
		Email:    payload.Email,
		Username: payload.Username,
		Hash:     hashedPassword,
		Salt:     salt,
		Name:     payload.Username,
	}
	return s.userRepo.Insert(err, userModel)
}
