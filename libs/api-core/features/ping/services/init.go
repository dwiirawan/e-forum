package ping

import (
	"gorm.io/gorm"
)

type PingService struct {
	db *gorm.DB
}

func NewPingService(db *gorm.DB) *PingService {
	return &PingService{
		db: db,
	}
}
