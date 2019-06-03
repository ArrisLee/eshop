package main

import (
	"eshop/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	routers := gin.Default()
	routers.GET("/api/v1/product", controller.CreateTestProduct)
	routers.GET("/api/v1/products", controller.GetAllProducts)
	routers.POST("/api/v1/login", controller.Login)
	routers.POST("/api/v1/register", controller.Register)
	routers.POST("/api/v1/order", controller.CreateOrder)
	routers.GET("/api/v1/orders", controller.CreateOrder)
	routers.POST("/api/v1/payment", controller.CreateOrder)
	routers.GET("/api/v1/payments", controller.CreateOrder)
	routers.Run(":8000")
}
