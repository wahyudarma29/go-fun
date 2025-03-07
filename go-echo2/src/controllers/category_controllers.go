// the controller is for handling http requests

package controllers

import (
	"go-echo2/src/models"
	"go-echo2/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func NewCategoryController(categoryRepo *service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryRepo}
}

func (s *CategoryController) GetAllCateories(c echo.Context) (error) {
	categories, err := s.CategoryService.GetAllCateories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch categories"})
	}
	return c.JSON(http.StatusOK, categories)
}

func (s *CategoryController) CreateCategory(c echo.Context) (error) {
	var category models.Category

	if err := c.Bind(&category); err != nil {
		return err
	}
	
	err := s.CategoryService.CreateCategory(&category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, category)
}
