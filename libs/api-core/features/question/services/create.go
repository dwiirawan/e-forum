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

	if len(payload.Tags) > 0 {
		questionTagModels := make([]models.QuestionTagsModel, len(payload.Tags))
		for i, tag := range payload.Tags {
			questionTagModels[i] = models.QuestionTagsModel{
				QuestionID: question.ID,
				TagID:      tag,
			}
		}
		if err := s.db.Create(&questionTagModels).Error; err != nil {
			return utils.NewError(fiber.StatusInternalServerError, "E_CREATE_QUESTION_TAG", "failed to create question tag", err)
		}
	}
	return nil
}
