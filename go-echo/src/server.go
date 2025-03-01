package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Product struct {
	ID string `json:"id"`
	ProductName string `json:"productName"`
	ProductDesc string `json:"productDesc"`
  }

type ReturnResponse struct {
	
}

var product []Product

func RemoveItem(s []Product, index int) []Product {
    return append(s[:index], s[index+1:]...)
}

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, product)
	})
	app.POST("/products", func(c echo.Context) error {
		newproduct := new(Product)
		if err := c.Bind(newproduct); err != nil {
			return err
		}
		for _, v := range product {
			if v.ID == newproduct.ID {
				return c.JSON(http.StatusBadRequest, "Data Already Exists.")
			}
		}

		product = append(product, *newproduct)
		return c.JSON(http.StatusCreated, newproduct)
	})
	app.PUT("/products/:id", func(c echo.Context) error {
		id := c.Param("id")
		updateproduct := new(Product)
		if err := c.Bind(updateproduct); err != nil {
			return err
		}
		for i, v := range product {
			if v.ID == id {
				product[i] = *updateproduct
				return c.JSON(http.StatusOK, product)
			} 
		}

		return c.String(http.StatusNotFound, "Data Not Found.")
	})
	app.DELETE("/products/:id", func(c echo.Context) error {
		id := c.Param("id")
		for i, v := range product {
			if v.ID == id {
				product = RemoveItem(product, i)
				return c.JSON(http.StatusOK, product)
			} 
		}

		return c.String(http.StatusNotFound, "Data Not Found.")
	})
	app.Logger.Fatal(app.Start(":1323"))
	if l, ok := app.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
}