// the controller is for handling http requests

package controllers

import (
	"go-echo2/src/models"
	"go-echo2/src/service"
	"net/http"
	"strconv"

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

func (s *CategoryController) GetCategoryByID(c echo.Context) (error) {
	id := c.Param("id")
	parseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ID not valid."})
	}

	categoryID := uint(parseID)

	category, err := s.CategoryService.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Data not found."})
	}
	return c.JSON(http.StatusOK, category)
}

func (s *CategoryController) EditCategory(c echo.Context) (error) {
	var category models.Category
	id := c.Param("id")
	parseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ID not valid."})
	}

	categoryID := uint(parseID)

	if err := c.Bind(&category); err != nil {
		return err
	}

	_, err = s.CategoryService.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}
	
	updatedCategory, err := s.CategoryService.EditCategory(&category, categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update category"})
	}

	return c.JSON(http.StatusOK, updatedCategory)
}

func (s *CategoryController) DeleteCategory(c echo.Context) (error) {
	id := c.Param("id")
	parseID, err := strconv.ParseUint(id, 10, 32)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ID not valid."})
	}

	categoryID := uint(parseID)

	err = s.CategoryService.DeleteCategory(categoryID)
	
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Data not found."})
	}
	return c.JSON(http.StatusNoContent, map[string]string{"message": "Category deleted."})
}