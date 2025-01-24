package models

type QuestionTags struct {
	QuestionID uint `json:"question_id" gorm:"column:question_id"`
	TagID      uint `json:"tag_id" gorm:"column:tag_id"`
}

func (QuestionTags) TableName() string {
	return "question_tags"
}
