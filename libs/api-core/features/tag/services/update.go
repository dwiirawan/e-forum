package services

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/tag/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
)

func (s *TagService) Update(id string, payload dto.TagUpdate) error {
	data := models.TagModel{
		Name: payload.Name,
	}
	err := s.db.Model(&data).Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_TAG", "failed to update tag", err)
	}
	return nil
}
