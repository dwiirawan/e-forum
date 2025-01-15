package services

import (
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *AnswerService) Delete(id string) error {
	if err := s.db.Delete(&models.AnswerModel{
		ID: id,
	}).Error; err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_DELETE_ANSWER", "failed to delete answer", err)
	}
	return nil
}
