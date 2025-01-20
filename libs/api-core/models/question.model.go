package models

import "time"

type QuestionModel struct {
	ID        string    `json:"id" gorm:"primary_key;column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Content   string    `json:"content" gorm:"type:text"`
	UserID    string    `json:"user_id" gorm:"type:uuid"`
	Views     int       `json:"views" gorm:"type:int"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp;default:null"`
	User      UserModel `json:"user" gorm:"foreignKey:ID;references:UserID"`
}

func (QuestionModel) TableName() string {
	return "questions"
}
