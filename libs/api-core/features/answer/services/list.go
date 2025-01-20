package services

import (
	"libs/api-core/features/answer/dto"
	userDto "libs/api-core/features/auth/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *AnswerService) List(questionId string) ([]dto.AnswerDetail, error) {
	var answers []models.AnswerModel
	if err := s.db.Model(&models.AnswerModel{}).Where("question_id = ?", questionId).Joins("User").Order("created_at DESC").Find(&answers).Error; err != nil {
		return nil, utils.NewError(fiber.StatusInternalServerError, "E_LIST_ANSWERS", "failed to list answers", err)
	}

	var data []dto.AnswerDetail

	for _, answer := range answers {
		data = append(data, dto.AnswerDetail{
			ID:         answer.ID,
			Content:    answer.Content,
			QuestionID: answer.QuestionID,
			User: userDto.UserIdentity{
				ID:       answer.UserID,
				Username: answer.User.Username,
				Email:    answer.User.Email,
				IsActive: answer.User.IsActive,
			},
			CreatedAt: answer.CreatedAt.Format(time.RFC3339),
			UpdatedAt: answer.UpdatedAt.Format(time.RFC3339),
		})
	}

	return data, nil
}
