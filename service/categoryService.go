package service

import (
	"example.com/m/v2/models"
	"example.com/m/v2/repo"
	"github.com/google/uuid"
)

type CategoryService struct {
	categoryRepo repo.ICategoryRepo
}

func NewCategoryService(categoryRepo repo.ICategoryRepo) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) FindAllCategories() ([]models.Category, error){
	return s.categoryRepo.FindAllCategories()
}

func (s *CategoryService) FindCategoryByID(categoryID uuid.UUID) (models.Category, error){
	return s.categoryRepo.FindCategoryByID(categoryID)
}

func (s *CategoryService) CreateCategory(category models.Category) error{
	return s.categoryRepo.CreateCategory(category)
}

func (s *CategoryService) UpdateCategoryByID(categoryID uuid.UUID, category models.Category) error{
	return s.categoryRepo.UpdateCategoryByID(categoryID, category)
}

func (s *CategoryService) DeleteCategoryByID(categoryID uuid.UUID) error{
	return s.categoryRepo.DeleteCategoryByID(categoryID)
}