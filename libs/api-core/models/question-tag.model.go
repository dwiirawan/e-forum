package models

type QuestionTagsModel struct {
	QuestionID string        `json:"question_id" gorm:"column:question_id"`
	TagID      string        `json:"tag_id" gorm:"column:tag_id"`
	Tag        TagModel      `json:"tag" gorm:"foreignKey:TagID;references:ID"`
	Question   QuestionModel `json:"question" gorm:"foreignKey:QuestionID;references:ID"`
}

func (QuestionTagsModel) TableName() string {
	return "question_tags"
}
