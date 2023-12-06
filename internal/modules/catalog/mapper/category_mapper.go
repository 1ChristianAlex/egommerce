package mapper

import (
	"khrix/egommerce/internal/core/addons"
	"khrix/egommerce/internal/modules/catalog/dto"
	"khrix/egommerce/internal/modules/catalog/repository/entities"

	"gorm.io/gorm"
)

type CategoryMapper struct{}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (m CategoryMapper) ToDto(item entities.Category) *dto.CategoryOutputDto {
	var subCategories []dto.CategoryOutputDto

	if len(item.SubCategory) > 0 {
		subCategories = addons.Map(item.SubCategory, func(item entities.Category) dto.CategoryOutputDto {
			if len(item.SubCategory) > 0 {
				return dto.CategoryOutputDto{
					ID:          item.ID,
					Name:        item.Name,
					SubCategory: addons.Map(item.SubCategory, func(item entities.Category) dto.CategoryOutputDto { return *m.ToDto(item) }),
				}
			}

			return dto.CategoryOutputDto{ID: item.ID, Name: item.Name}
		})
	}

	return &dto.CategoryOutputDto{ID: item.ID, Name: item.Name, SubCategory: subCategories}
}

func (m CategoryMapper) ToEntity(item dto.CategoryOutputDto) *entities.Category {
	var subCategories []entities.Category

	if len(item.SubCategory) > 0 {
		subCategories = addons.Map(item.SubCategory, func(item dto.CategoryOutputDto) entities.Category {
			if len(item.SubCategory) > 0 {
				subCateg := addons.Map(item.SubCategory, func(item dto.CategoryOutputDto) entities.Category { return *m.ToEntity(item) })

				return entities.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name, SubCategory: subCateg}
			}

			return entities.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name}
		})
	}
	return &entities.Category{Model: gorm.Model{ID: item.ID}, Name: item.Name, SubCategory: subCategories}
}
