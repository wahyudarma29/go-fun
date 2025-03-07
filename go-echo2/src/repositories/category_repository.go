// repositories is the place for the database operations

package repositories

import (
	"go-echo2/src/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAllCateories() ([]models.Category, error) {
	var categories []models.Category
	result := r.DB.Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) CreateCategory(category *models.Category) (error) {
	return r.DB.Create(category).Error
}
