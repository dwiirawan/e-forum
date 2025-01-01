package services

import (
	"gorm.io/gorm"
)

type QuestionService struct {
	db *gorm.DB
}

func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{
		db: db,
	}
}
