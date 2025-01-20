package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/comment/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
)

func (s *CommentService) Create(payload dto.CommentCreate, userId string) (*string, error) {
	fmt.Println(payload.ParentType, "<<<< IKIII")
	comment := models.CommentModel{
		Content:    payload.Content,
		ParentType: payload.ParentType,
		ParentID:   payload.ParentID,
		UserID:     userId,
	}

	if err := s.db.Create(&comment).Error; err != nil {
		return nil, utils.NewError(fiber.StatusInternalServerError, "E_CREATE_COMMENT", "failed to create comment", err)
	}

	return &comment.ID, nil
}
