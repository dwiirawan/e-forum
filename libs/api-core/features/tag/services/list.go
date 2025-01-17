package services

import (
	"github.com/gofiber/fiber/v2"
	"libs/api-core/features/tag/dto"
	"libs/api-core/utils"

	"libs/api-core/models"
)

func (s *TagService) List() (list []dto.Tag, err error) {
	var raw []models.TagModel

	err = s.db.Order("name ASC").Find(&raw).Error
	if err != nil {
		return nil, utils.NewError(fiber.StatusNotFound, "E_LIST_TAG", "failed to list tag", err)
	}

	var data []dto.Tag

	for _, v := range raw {
		data = append(data, dto.Tag{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return data, nil

}
