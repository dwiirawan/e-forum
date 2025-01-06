package models

import (
	"time"
)

type AnswerModel struct {
	ID         string        `json:"id" gorm:"primary_key;column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	QuestionID string        `json:"question_id" gorm:"type:uuid"`
	UserID     string        `json:"user_id" gorm:"type:uuid"`
	Content    string        `json:"content" gorm:"type:text"`
	Votes      int           `json:"votes" gorm:"type:int;default:0"`
	CreatedAt  time.Time     `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time     `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
	User       UserModel     `json:"user" gorm:"foreignKey:ID;references:UserID"`
	Question   QuestionModel `json:"question" gorm:"foreignKey:ID;references:QuestionID"`
}

func (AnswerModel) TableName() string {
	return "answers"
}
