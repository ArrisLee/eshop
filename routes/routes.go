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
	e.POST("/api/v1/user", controller.Register)
	e.GET("/api/v1/products", controller.GetProducts)
	// Start server
	e.Logger.Fatal(e.Start(":8500"))
}
