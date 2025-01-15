package services

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/comment/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
	"time"
)

func (s *CommentService) List(ParentID string) ([]dto.Comment, error) {
	var comments []models.CommentModel
	if err := s.db.Model(&models.CommentModel{}).Where("parent_id = ?", ParentID).Order("created_at ASC").Find(&comments).Error; err != nil {
		return nil, utils.NewError(fiber.StatusInternalServerError, "E_LIST_COMMENT", "failed to list comment", err)
	}

	var data []dto.Comment

	for _, comment := range comments {
		data = append(data, dto.Comment{
			ID:         comment.ID,
			Content:    comment.Content,
			ParentType: comment.ParentType,
			ParentID:   comment.ParentID,
			UserID:     comment.UserID,
			CreatedAt:  comment.CreatedAt.Format(time.RFC3339),
		})
	}

	return data, nil
}
