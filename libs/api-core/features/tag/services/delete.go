package services

import (
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *TagService) Delete(id string) error {

	err := s.db.Delete(&models.TagModel{ID: id}).Error
	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_DELETE_TAG", "failed to delete tag", err)
	}

	return nil
}
