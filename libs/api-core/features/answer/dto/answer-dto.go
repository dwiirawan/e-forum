package dto

import auth "libs/api-core/features/auth/dto"

type AnswerCreate struct {
	Content    string `json:"content"`
	QuestionID string `json:"question_id"`
}

type AnswerUpdate struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	QuestionID string `json:"question_id"`
}

type AnswerDetail struct {
	ID         string            `json:"id"`
	Content    string            `json:"content"`
	QuestionID string            `json:"question_id"`
	User       auth.UserIdentity `json:"user"`
	CreatedAt  string            `json:"created_at"`
	UpdatedAt  string            `json:"updated_at"`
}
