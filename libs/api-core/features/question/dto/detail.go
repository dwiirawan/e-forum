package dto

import (
	auth "libs/api-core/features/auth/dto"
	"time"
)

type QuestionDetail struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	User      auth.UserIdentity `json:"user"`
}
