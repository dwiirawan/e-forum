package dto

type CreateQuestionTagDto struct {
	QuestionID uint `json:"question_id" validate:"required"`
	TagID      uint `json:"tag_id" validate:"required"`
}
