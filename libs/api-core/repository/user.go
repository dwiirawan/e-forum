package repository

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"libs/api-core/models"
	"libs/api-core/utils"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindByEmailOrUsername(username string) (*models.UserModel, error) {
	var user models.UserModel
	if err := r.DB.Where("email = ? OR username = ?", username, username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewError(fiber.StatusNotFound, "E_FIND_USER_NOT_FOUND", "user not found", err)
		}
		return nil, utils.NewError(fiber.StatusInternalServerError, "E_FIND_USER_FAILED", utils.ERR_INTERNAL_SERVER_ERROR, err)
	}
	return &user, nil
}
func (r *UserRepository) Insert(err error, userModel models.UserModel) error {
	err = r.DB.Model(userModel).Create(&userModel).Error

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
