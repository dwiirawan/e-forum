package services

import (
	"gorm.io/gorm"
)

type VoteService struct {
	db *gorm.DB
}

func NewVoteService(db *gorm.DB) *VoteService {
	return &VoteService{
		db: db,
	}
}
