package services

import (
	"libs/api-core/features/question/dto"
	"libs/api-core/models"
	"libs/api-core/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s *QuestionService) Update(id string, payload dto.QuestionUpdate) error {
	question := models.QuestionModel{
		Title:     payload.Title,
		Content:   payload.Content,
		UpdatedAt: time.Now(),
	}
	err := s.db.Model(&question).Where("id = ?", id).Updates(payload).Error
	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_QUESTION", "failed to update question", err)
	}
	return nil
}

func (s *QuestionService) UpdateTags(id string, tags []string) error {
	oldTags := []models.QuestionTagsModel{}

	err := s.db.Where("question_id = ?", id).Find(&oldTags).Error
	if err != nil {
		return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_QUESTION", "failed to update question", err)
	}
	for _, tag := range oldTags {
		err := s.db.Delete(&tag).Error
		if err != nil {
			return utils.NewError(fiber.StatusInternalServerError, "E_UPDATE_QUESTION", "failed to update question", err)
		}
	}

	updateTag := []models.QuestionTagsModel{}

	for _, tag := range tags {
		updateTag = append(updateTag, models.QuestionTagsModel{
			QuestionID: id,
			TagID:      tag,
		})
	}

	return nil

}
