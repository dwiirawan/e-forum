package services

import (
	"libs/api-core/features/answer/dto"
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *AnswerService) Create(payload dto.AnswerCreate, userId string) (*string, error) {
	answer := models.AnswerModel{
		QuestionID: payload.QuestionID,
		UserID:     userId,
		Content:    payload.Content,
	}
	if err := s.db.Create(&answer).Error; err != nil {
		return nil, utils.NewError(fiber.StatusInternalServerError, "E_CREATE_QUESTION", "failed to create question", err)
	}

	return &answer.ID, nil
}
