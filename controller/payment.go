package controller

import (
	"eshop/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo"
)

//MakePayment func
func MakePayment(c echo.Context) error {
	if !Authorize(c) {
		return c.JSON(http.StatusUnauthorized, "forbidden")
	}
	type paymentRequest struct {
		PaymentID primitive.ObjectID `json:"paymentID"`
		Source    string             `json:"source"`
	}
	req := &paymentRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//call stripe here
	result := true
	if err := db.UpdatePaymentResult(req.PaymentID, result); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}

//GetPayments func
func GetPayments(c echo.Context) error {
	if !Authorize(c) {
		return c.JSON(http.StatusUnauthorized, "forbidden")
	}
	customerID, _ := primitive.ObjectIDFromHex(c.Request().Header.Get("id"))
	payments, err := db.ReadPaymentsByCustomerID(customerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payments)
}
