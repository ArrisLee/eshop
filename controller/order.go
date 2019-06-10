package controller

import (
	"crypto/md5"
	"eshop/db"
	"fmt"
	"net/http"
	"strconv"

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
	order.ShortID = shortID(order.ID)
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

func shortID(orderID primitive.ObjectID) string {
	s := orderID.Hex()
	sMd5 := md5.Sum([]byte(s))
	inputSlice := []rune(fmt.Sprintf("%x", sMd5)[0:24])
	var inputInt []int64
	for i := 0; i < len(s); i = i + 4 {
		n, _ := strconv.ParseInt(string(inputSlice[i:i+4]), 16, 64)
		inputInt = append(inputInt, n)
	}
	for i, v := range inputInt {
		inputInt[i] = remainder(v)
	}
	table := "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	var result string
	for i := len(inputInt) - 1; i >= 0; i-- {
		result = result + fmt.Sprintf("%c", table[inputInt[i]])
	}

	return result
}

func remainder(n int64) int64 {
	if n < 32 {
		return n
	}
	var totalRemainder int64
	for n > 0 {
		totalRemainder = totalRemainder + n%32
		n = n / 32
	}
	return remainder(totalRemainder)
}
