package models

import "time"

type CommentModel struct {
	ID         string    `json:"id" gorm:"primary_key;column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Content    string    `json:"content" gorm:"column:content;type:text;not null"`
	ParentType int       `json:"parent_type" gorm:"column:parent_type;type:smallint;not null"`
	ParentID   string    `json:"parent_id" gorm:"column:parent_id;type:uuid;not null"`
	UserID     string    `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
}

func (CommentModel) TableName() string {
	return "comments"
}
