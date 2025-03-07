package main

import (
	"go-echo2/src/controllers"
	"go-echo2/src/db"
	"go-echo2/src/repositories"
	"go-echo2/src/routes"
	"go-echo2/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)


func main() {
	db.Connect()
	// Initialize repositories, services, and controllers
	categoryRepo := repositories.NewCategoryRepository(db.DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)
	
	app := echo.New()
	routes.RegisterRoutes(app, categoryController)
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.Logger.Fatal(app.Start(":1323"))
	if l, ok := app.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
}