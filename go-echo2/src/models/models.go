package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  		string    	`json:"productName"`
	ProductDescription  string    	`json:"productDescription"`
	Quantity			int32		`json:"quantity"`
	CategoryID			uint		`json:"categoryID"`
}

type Category struct {
	gorm.Model
	CategoryName  			string  `json:"categoryName"`
	CategoryDescription  	string  `json:"categoryDescription"`
	Products 				[]Product
}


