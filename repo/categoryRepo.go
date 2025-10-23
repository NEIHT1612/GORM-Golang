package repo

import (
	"example.com/m/v2/models"
	"github.com/google/uuid"
)

type ICategoryRepo interface {
	FindAllCategories() ([]models.Category, error)
	FindCategoryByID(categoryID uuid.UUID) (models.Category, error)
	CreateCategory(category models.Category) error
	UpdateCategoryByID(categoryID uuid.UUID, category models.Category) error
	DeleteCategoryByID(categoryID uuid.UUID) error
}

type categoryRepo struct{}

// DeleteCategoryByID implements ICategoryRepo.
func (c *categoryRepo) DeleteCategoryByID(categoryID uuid.UUID) error {
	var cateModel models.Category
	return cateModel.DeleteCategoryByID(categoryID)
}

// UpdateCategoryByID implements ICategoryRepo.
func (c *categoryRepo) UpdateCategoryByID(categoryID uuid.UUID, category models.Category) error {
	var cateModel models.Category
	return cateModel.UpdateCategoryByID(categoryID, category)
}

// CreateCategory implements ICategoryRepo.
func (c *categoryRepo) CreateCategory(category models.Category) error {
	var cateModel models.Category
	return cateModel.CreateCategory(category)
}

// FindCategoryByID implements ICategoryRepo.
func (c *categoryRepo) FindCategoryByID(categoryID uuid.UUID) (models.Category, error) {
	var cateModel models.Category
	return cateModel.FindCategoryByID(categoryID)
}

// FindAllCategories implements ICategoryRepo.
func (c *categoryRepo) FindAllCategories() ([]models.Category, error) {
	var cateModel models.Category
	return cateModel.FindAllCategories()
}

func NewCategoryRepo() ICategoryRepo {
	return &categoryRepo{}
}
