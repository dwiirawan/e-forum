package models

type VoteType int
type ParentType int

const (
	VoteTypeUp   VoteType = 1
	VoteTypeDown VoteType = 2

	ParentTypeQuestion ParentType = 1
	ParentTypeAnswer   ParentType = 2
	ParentTypeComment  ParentType = 3
)

type VoteModel struct {
	ID         string     `gorm:"primaryKey;column:id;type:uuid;default:gen_random_uuid()"`
	ParentType ParentType `gorm:"column:parent_type;type:smallint"`
	ParentID   string     `gorm:"column:parent_id;type:uuid"`
	UserID     string     `gorm:"column:user_id;type:uuid"`
	VoteType   VoteType   `gorm:"column:vote_type;type:smallint"`
}

func (VoteModel) TableName() string {
	return "votes"
}
