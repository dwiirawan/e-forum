package services

import (
	"libs/api-core/features/answer/dto"
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *AnswerService) Update(id string, payload dto.AnswerUpdate) error {
	answer := models.AnswerModel{
		Content: payload.Content,
	}
	if err := s.db.Model(&answer).Where("id = ?", id).Updates(answer).Error; err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_ANSWER", "failed to update answer", err)
	}
	return nil
}
