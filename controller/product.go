package controller

import (
	"eshop/db"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTestProduct(c echo.Context) error {
	product := &db.Product{}
	product.ID = primitive.NewObjectID()
	product.Name = "Apples"
	product.Img = "https://a.fsimg.co.nz/product/retail/fan/image/400x400/5046525.png"
	product.Price = 2.99
	product.Desc = "Your everyday fresh vegetables and fruits delivery service."
	if _, err := db.AddProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, product)
}

func GetProducts(c echo.Context) error {
	products, err := db.ReadAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}
