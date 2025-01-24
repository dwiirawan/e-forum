package services

import (
	"gorm.io/gorm"
)

type QuestionTagService struct {
	db *gorm.DB
}

func NewQuestionTagService(db *gorm.DB) *QuestionTagService {
	return &QuestionTagService{
		db: db,
	}
}
