package controller

import (
	"eshop/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTestProduct(c *gin.Context) {
	product := &db.Product{}
	product.ID = primitive.NewObjectID()
	product.Name = "Switch"
	product.Pic = "https://www.pbtech.co.nz/imgprod/G/A/GAMNTD1003__1.jpg"
	product.Price = 499.9
	if _, err := db.AddProduct(product); err != nil {
		c.JSON(500, gin.H{"data": err.Error()})
		return
	}
	c.JSON(201, gin.H{"data": product})
}

func GetAllProducts(c *gin.Context) {
	// product := db.Product{}
	// product.ID = "5caef7f7042eb13e827ec928"
	// product.Name = "Switch"
	// product.Pic = "https://www.pbtech.co.nz/imgprod/G/A/GAMNTD1003__1.jpg"
	// product.Price = 499.9
	// productTwo := db.Product{}
	// productTwo.ID = "5caef7f7042eb13e827ec927"
	// productTwo.Name = "Play Station 4"
	// productTwo.Pic = "https://www.pbtech.co.nz/imgprod/G/A/GAMSNY4703__1.jpg"
	// productTwo.Price = 588.8
	// productThree := db.Product{}
	// productThree.ID = "5caef7f7042eb13e827ec926"
	// productThree.Name = "Xbox ONE"
	// productThree.Pic = "https://www.pbtech.co.nz/thumbs/G/A/GAMMST14014.jpg.large.jpg"
	// productThree.Price = 521.9
	// prodcuts := []db.Product{product, productTwo, productThree}
	products, err := db.ReadAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"data": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": products})
}
