package services

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/question/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
)

func (s *QuestionService) Create(payload dto.QuestionCreate, userId string) error {
	question := models.QuestionModel{
		Content: payload.Content,
		Title:   payload.Title,
		UserID:  userId,
	}

	if err := s.db.Create(&question).Error; err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_CREATE_QUESTION", "failed to create question", err)
	}
	return nil
}
