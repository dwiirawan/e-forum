package dto

type DeleteQuestionTagDto struct {
	QuestionID string `json:"question_id" validate:"required"`
	TagID      string `json:"tag_id" validate:"required"`
}
