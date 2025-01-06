package services

import (
	"libs/api-core/features/answer/dto"
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *AnswerService) Get(id string) (*dto.AnswerDetail, error) {
	var answer dto.AnswerDetail
	if err := s.db.Model(&models.AnswerModel{}).Where("id = ?", id).Joins("User").First(&answer).Error; err != nil {
		return nil, utils.NewError(fiber.StatusInternalServerError, "E_GET_ANSWER", "failed to get answer", err)
	}
	return &answer, nil
}
