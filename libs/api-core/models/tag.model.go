package models

type TagModel struct {
	ID   string `json:"id" gorm:"primary_key;column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string `json:"name" gorm:"column:name;type:varchar(255);not null"`
}

func (TagModel) TableName() string {
	return "tags"
}
