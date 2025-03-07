// repositories is the place for the database operations

package repositories

import (
	"go-echo2/src/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *CategoryRepository) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	result := r.DB.First(&category, id)
	return category, result.Error
}

func (r *CategoryRepository) CreateCategory(category *models.Category) (error) {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) EditCategory(category *models.Category, id uint) (models.Category, error) {
	var updatedCategory models.Category
	result := r.DB.Model(&models.Category{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(category).Scan(&updatedCategory)
	return updatedCategory, result.Error
}	

func (r *CategoryRepository) DeleteCategory(category *models.Category) (error) {
	result := r.DB.Model(&models.Category{}).Delete(category)
	return result.Error
}
