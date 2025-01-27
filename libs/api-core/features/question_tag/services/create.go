package services

import (
	"libs/api-core/features/question_tag/dto"
	"libs/api-core/models"
)

func (a *QuestionTagService) Create(payload dto.CreateQuestionTagDto) error {
	questionTag := models.QuestionTagsModel{
		QuestionID: payload.QuestionID,
		TagID:      payload.TagID,
	}
	err := a.db.Model(questionTag).Create(&questionTag).Error
	if err != nil {
		return err
	}

	return nil
}
