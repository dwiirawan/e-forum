package models

type QuestionModel struct {
	ID        string    `json:"id" gorm:"primary_key;column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Content   string    `json:"content" gorm:"type:text"`
	UserID    string    `json:"user_id" gorm:"type:uuid"`
	Views     int       `json:"views" gorm:"type:int"`
	CreatedAt string    `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt string    `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt string    `json:"deleted_at" gorm:"type:timestamp"`
	User      UserModel `json:"user" gorm:"foreignKey:ID;references:UserID"`
}

func (QuestionModel) TableName() string {
	return "questions"
}
