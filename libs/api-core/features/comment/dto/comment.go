package dto

type Comment struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	ParentType int    `json:"parent_type"`
	ParentID   string `json:"parent_id"`
	UserID     string `json:"user_id"`
	CreatedAt  string `json:"created_at"`
}

type CommentCreate struct {
	Content    string `json:"content"`
	ParentType int    `json:"parent_type"`
	ParentID   string `json:"parent_id"`
}
