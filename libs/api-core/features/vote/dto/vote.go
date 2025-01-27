package dto

import "libs/api-core/models"

type Vote struct {
	ParentType models.ParentType `json:"parent_type"`
	ParentID   string            `json:"parent_id"`
}

type UnVote struct {
	ParentType models.ParentType `json:"parent_type"`
	ParentID   string            `json:"parent_id"`
}
