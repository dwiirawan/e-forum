package dto

type DeleteQuestionTagDto struct {
	QuestionID uint `json:"question_id" validate:"required"`
	TagID      uint `json:"tag_id" validate:"required"`
}
