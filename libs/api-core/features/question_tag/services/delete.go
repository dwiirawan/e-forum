package services

import (
	"libs/api-core/features/question_tag/dto"
	"libs/api-core/models"
)

func (a *QuestionTagService) Delete(questionTag dto.DeleteQuestionTagDto) error {
	model := models.QuestionTagsModel{
		QuestionID: questionTag.QuestionID,
		TagID:      questionTag.TagID,
	}
	err := a.db.Model(&model).Delete(model).Error
	if err != nil {
		return err
	}

	return nil
}
