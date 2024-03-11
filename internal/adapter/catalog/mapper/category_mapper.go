package mapper

import (
	"khrix/egommerce/internal/application/catalog/dto"
	"khrix/egommerce/internal/application/libs/addons"
	"khrix/egommerce/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (m CategoryMapper) ToDto(item models.Category) *dto.CategoryOutputDto {
	var subCategories []dto.CategoryOutputDto

	if len(item.SubCategory) > 0 {
		subCategories = addons.Map(item.SubCategory, func(item models.Category) dto.CategoryOutputDto {
			if len(item.SubCategory) > 0 {
				return dto.CategoryOutputDto{
					ID:          item.ID,
					Name:        item.Name,
					SubCategory: addons.Map(item.SubCategory, func(item models.Category) dto.CategoryOutputDto { return *m.ToDto(item) }),
				}
			}

			return dto.CategoryOutputDto{ID: item.ID, Name: item.Name}
		})
	}

	return &dto.CategoryOutputDto{ID: item.ID, Name: item.Name, SubCategory: subCategories}
}

func (m CategoryMapper) ToEntity(item dto.CategoryOutputDto) *models.Category {
	var subCategories []models.Category

	if len(item.SubCategory) > 0 {
		subCategories = addons.Map(item.SubCategory, func(item dto.CategoryOutputDto) models.Category {
			if len(item.SubCategory) > 0 {
				subCateg := addons.Map(item.SubCategory, func(item dto.CategoryOutputDto) models.Category { return *m.ToEntity(item) })

				return models.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name, SubCategory: subCateg}
			}

			return models.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name}
		})
	}
	return &models.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name, SubCategory: subCategories}
}
