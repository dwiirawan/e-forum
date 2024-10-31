package user

import "libs/api-core/models"

func (a *Service) GetUser(ID string) (models.UserModel, error) {
	var user models.UserModel
	err := a.db.Where("id = ?", ID).First(&user).Error
	return user, err
}
