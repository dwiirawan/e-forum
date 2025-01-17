package dto

type TagCreate struct {
	Name string `json:"name"`
}

type TagUpdate struct {
	Name string `json:"name"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
