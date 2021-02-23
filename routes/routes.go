package routes

import (
	"github.com/ArrisLee/Eshop/controller"

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
	e.POST("/api/v1/cart", controller.CreateCart)
	e.POST("/api/v1/order", controller.CreateOrder)
	e.POST("/api/v1/payment", controller.MakePayment)
	e.GET("/api/v1/products", controller.GetProducts)
	e.GET("/api/v1/payments", controller.GetPayments)
	e.GET("/api/v1/orders", controller.GetOrders)
	e.GET("/api/v1/email", controller.SendEmailNotification)
	// Start server
	e.Logger.Fatal(e.Start(":8500"))
}
