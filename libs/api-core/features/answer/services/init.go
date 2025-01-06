package services

import (
	"gorm.io/gorm"
)

type AnswerService struct {
	db *gorm.DB
}

func NewAnswerService(db *gorm.DB) *AnswerService {
	return &AnswerService{
		db: db,
	}
}
