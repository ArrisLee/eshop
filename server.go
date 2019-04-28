package main

import (
	"eshop/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()
	routers.GET("/api/v1/product", controller.CreateTestProduct)
	routers.GET("/api/v1/products", controller.GetAllProducts)
	routers.GET("/api/v1/customers", controller.GetAllCustomers)
	routers.POST("/api/v1/order", controller.CreateOrder)
	routers.Run(":8000")
}
