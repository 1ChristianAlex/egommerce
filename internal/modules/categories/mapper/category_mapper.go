package mapper

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/categories/dto"
	"khrix/egommerce/internal/modules/categories/repository/entities"

	"gorm.io/gorm"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (m CategoryMapper) ToDto(item entities.Category) dto.CategoryOutputDto {
	var subCategories []dto.CategoryOutputDto

	if len(item.Category) > 0 {
		subCategories = addons.Map(item.Category, func(item entities.Category) dto.CategoryOutputDto {
			if len(item.Category) > 0 {
				return dto.CategoryOutputDto{ID: item.ID, Name: item.Name, SubCategory: addons.Map(item.Category, m.ToDto)}
			}

			return dto.CategoryOutputDto{ID: item.ID, Name: item.Name}
		})
	}

	return dto.CategoryOutputDto{ID: item.ID, Name: item.Name, SubCategory: subCategories}
}

func (m CategoryMapper) ToEntity(item dto.CategoryOutputDto) entities.Category {
	var subCategories []entities.Category

	if len(item.SubCategory) > 0 {
		subCategories = addons.Map(item.SubCategory, func(item dto.CategoryOutputDto) entities.Category {
			if len(item.SubCategory) > 0 {
				return entities.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name, Category: addons.Map(item.SubCategory, m.ToEntity)}
			}

			return entities.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name}
		})
	}
	return entities.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name, Category: subCategories}
}
