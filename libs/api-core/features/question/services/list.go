package services

import (
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/features/question/dto"
	tagDto "libs/api-core/features/tag/dto"
	"libs/api-core/models"
	"libs/api-core/utils"

	"github.com/gofiber/fiber/v2"
)

func (s *QuestionService) List(req *dto.PaginationListQuestionRequest) ([]dto.QuestionDetail, *dto.PaginationListQuestionMeta, error) {
	var questions []models.QuestionModel
	var totalData int64
	offset := (req.Page - 1) * req.Limit
	query := s.db.Model(&models.QuestionModel{}).Joins("User").Preload("QuestionTags").Preload("QuestionTags.Tag").Count(&totalData)

	if req.Search != "" {
		query = query.Where("title LIKE ?", "%"+req.Search+"%")
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order(req.Sort + " " + req.Order)
	}

	query = query.Limit(req.Limit).Offset(offset).Order("created_at desc")
	if err := query.Find(&questions).Error; err != nil {
		return nil, nil, utils.NewError(fiber.StatusInternalServerError, "E_LIST_QUESTION", "failed to list question", err)
	}
	totalPage := int(totalData) / req.Limit
	if int(totalData)%req.Limit > 0 {
		totalPage++
	}

	var data []dto.QuestionDetail

	for _, question := range questions {

		tags := make([]tagDto.Tag, len(question.QuestionTags))
		for i, tag := range question.QuestionTags {
			tags[i] = tagDto.Tag{ID: tag.Tag.ID, Name: tag.Tag.Name}
		}

		data = append(data, dto.QuestionDetail{
			ID:        question.ID,
			Title:     question.Title,
			Content:   question.Content,
			CreatedAt: question.CreatedAt,
			UpdatedAt: question.UpdatedAt,
			Votes:     question.Votes,
			Tags:      tags,
			User: auth.UserIdentity{
				ID:       question.User.ID.String(),
				Username: question.User.Username,
				Email:    question.User.Email,
				IsActive: question.User.IsActive,
			},
		})
	}

	return data, &dto.PaginationListQuestionMeta{
		CurrentPage: req.Page,
		TotalPage:   totalPage,
		TotalData:   int(totalData),
	}, nil

}
