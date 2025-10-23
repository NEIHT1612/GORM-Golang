package models

import (
	"example.com/m/v2/config"
	"github.com/google/uuid"
)

type Category struct {
	CategoryID   uuid.UUID `json:"category_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CategoryName string    `json:"category_name"`
}

// Mapping to TableName
func (Category) TableName() string {
	return "categories"
}

func (c *Category) FindAllCategories() ([]Category, error) {
	var categories []Category
	result := config.DB.Find((&categories))
	return categories, result.Error
}

func (c *Category) FindCategoryByID(categoryID uuid.UUID) (Category, error) {
	var category Category
	result := config.DB.First((&category), categoryID)
	return category, result.Error
}

func (c *Category) CreateCategory(category Category) error {
	result := config.DB.Create((&category))
	return result.Error
}

func (c *Category) UpdateCategoryByID(categoryID uuid.UUID, category Category) error{
	result := config.DB.Model(&Category{}).
						Where("category_id = ?", categoryID).
						Updates(map[string]interface{}{
							"category_name": category.CategoryName,
						})
	return result.Error
}

func (c *Category) DeleteCategoryByID(categoryID uuid.UUID) error{
	result := config.DB.Delete(&Category{}, "category_id = ?", categoryID)
	return result.Error
}