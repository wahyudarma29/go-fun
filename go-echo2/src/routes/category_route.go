package routes

import (
	Controller "go-echo2/src/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, categoryController *Controller.CategoryController) {
	e.GET("/categories", categoryController.GetAllCateories)
	e.POST("/categories", categoryController.CreateCategory)
}