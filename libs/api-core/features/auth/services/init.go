package auth

import (
	"gorm.io/gorm"
	"libs/api-core/common"
	"libs/api-core/repository"
)

type AuthService struct {
	db            *gorm.DB
	userRepo      *repository.UserRepository
	hashingMethod *common.HashingMethod
	jwtAuth       *common.BasicJwtAuth
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db:            db,
		userRepo:      repository.NewUserRepository(db),
		hashingMethod: common.NewHashingMethod(),
		jwtAuth:       common.NewBasicJwtAuth(),
	}
}
