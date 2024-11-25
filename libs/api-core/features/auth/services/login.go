package auth

import (
	"github.com/gofiber/fiber/v2"
	dto "libs/api-core/features/auth/dto"
	"libs/api-core/utils"
)

func (s *AuthService) Login(payload *dto.LoginRequestDTO) (res *dto.LoginResponseDTO, err error) {
	user, err := s.userRepo.FindByEmailOrUsername(payload.Username)
	if err != nil {
		return nil, err
	}

	valid := s.hashingMethod.VerifyPassword(payload.Password, user.Hash, user.Salt)
	if !valid {
		return nil, utils.NewError(fiber.StatusUnauthorized, "E_UNAUTHORIZED", "invalid username or password", nil)
	}

	userIdentity := &dto.UserIdentity{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		IsActive: user.IsActive,
	}

	token, err := s.jwtAuth.GenerateToken(userIdentity)

	if err != nil {
		return nil, err
	}
	return &dto.LoginResponseDTO{
		Token: token,
	}, nil

}
