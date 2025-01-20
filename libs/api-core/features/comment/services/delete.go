package services

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/models"
	"libs/api-core/utils"
)

func (s *CommentService) Delete(ID string) error {
	err := s.db.Delete(&models.CommentModel{
		ID: ID,
	}).Error

	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_DELETE_COMMENT", "failed to delete comment", err)
	}
	return nil
}
