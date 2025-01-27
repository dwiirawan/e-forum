package dto

import (
	auth "libs/api-core/features/auth/dto"
	"libs/api-core/features/tag/dto"
	"time"
)

type QuestionDetail struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	User      auth.UserIdentity `json:"user"`
	Tags      []dto.Tag         `json:"tags"`
	Votes     uint              `json:"votes"`
}
