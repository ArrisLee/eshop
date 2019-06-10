package controller

import (
	"eshop/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo"
)

//CreateOrder func
func CreateOrder(c echo.Context) error {
	if !Authorize(c) {
		return c.JSON(http.StatusUnauthorized, "forbidden")
	}
	type orderRequest struct {
		CartID  primitive.ObjectID `json:"cartID"`
		Name    string             `json:"name"`
		Phone   string             `json:"phone"`
		Address string             `json:"address"`
	}
	req := &orderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	cart, err := db.ReadCart(req.CartID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//init payment struct here
	order := &db.Order{}
	order.ID = primitive.NewObjectID()
	order.CustomerID, _ = primitive.ObjectIDFromHex(c.Request().Header.Get("id"))
	order.PaymentID = primitive.NewObjectID() //modify later
	order.Cart = cart
	order.Price = cart.TotalPrice
	order.Phone = req.Phone
	order.Address = req.Address
	order.ShortID = "TESTX" //modify later
	order.Paid = false
	order.Dispatched = false
	order.Delivered = false
	if err := db.AddOrder(order); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, order.ShortID)
}

//GetOrders func
func GetOrders(c echo.Context) error {
	if !Authorize(c) {
		return c.JSON(http.StatusUnauthorized, "forbidden")
	}
	customerID, _ := primitive.ObjectIDFromHex(c.Request().Header.Get("id"))
	orders, err := db.ReadOrdersByCustomerID(customerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}
