package models

import (
	"github.com/google/uuid"
	"time"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255);" json:"name"`
	Username  string    `gorm:"column:username;type:varchar(255);unique" json:"username"`
	Email     string    `gorm:"column:email;type:varchar(255);unique" json:"email"`
	Hash      string    `gorm:"column:hash;type:varchar(255);" json:"hash"`
	Salt      string    `gorm:"column:salt;type:varchar(255);" json:"salt"`
	IsActive  bool      `gorm:"column:is_active;type:boolean;" json:"is_active"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
