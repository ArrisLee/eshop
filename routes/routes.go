package routes

import (
	"eshop/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Run func
func Run() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	// Routes
	e.POST("/api/v1/user/register", controller.Register)
	e.POST("/api/v1/user/login", controller.Login)
	e.GET("/api/v1/products", controller.GetProducts)
	e.POST("/api/v1/cart", controller.CreateCart)
	e.POST("/api/v1/order", controller.CreateOrder)
	e.GET("/api/v1/email", controller.SendEmailNotification)
	// Start server
	e.Logger.Fatal(e.Start(":8500"))
}
