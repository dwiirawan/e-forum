package services

import (
	"libs/api-core/features/tag/dto"
	"libs/api-core/models"
)

func (s *TagService) Create(payload dto.TagCreate) (string, error) {
	tag := models.TagModel{
		Name: payload.Name,
	}

	if err := s.db.Create(&tag).Error; err != nil {
		return "", err
	}

	return tag.ID, nil

}
