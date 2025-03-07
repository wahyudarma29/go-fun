// service is place for the business logic
package service

import (
	"go-echo2/src/models"
	"go-echo2/src/repositories"
)

type CategoryService struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: categoryRepo}
}

func (s *CategoryService) GetAllCateories() ([]models.Category, error) {
	return s.CategoryRepo.GetAllCateories()
}

func (s *CategoryService) CreateCategory(category *models.Category) (error) {
	return s.CategoryRepo.CreateCategory(category)
}
