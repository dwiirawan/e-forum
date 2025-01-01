package services

import (
	"errors"
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/features/question/dto"
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (s *QuestionService) Get(id string) (*dto.QuestionDetail, error) {
	question := models.QuestionModel{
		ID: id,
	}

	if err := s.db.Joins("User").First(&question, question).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewError(fiber.StatusNotFound, "E_QUESTION_NOT_FOUND", "question not found", err)
		}

		return nil, utils.NewError(fiber.StatusInternalServerError, "E_GET_QUESTION", "failed to get detail question", err)
	}

	return &dto.QuestionDetail{
		ID:        question.ID,
		Title:     question.Title,
		Content:   question.Content,
		CreatedAt: question.CreatedAt,
		UpdatedAt: question.UpdatedAt,
		User: auth.UserIdentity{
			ID:       question.User.ID.String(),
			Username: question.User.Username,
			Email:    question.User.Email,
			IsActive: question.User.IsActive,
		},
	}, nil
}
