package services

import (
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *QuestionService) Delete(id string) error {

	err := s.db.Delete(&models.QuestionModel{ID: id}).Error
	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_DELETE_QUESTION", "failed to delete question", err)
	}

	return nil
}
