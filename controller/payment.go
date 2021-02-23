package controller

import (
	"net/http"

	"github.com/ArrisLee/Eshop/db"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
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
	payment, err := db.ReadPayment(req.PaymentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//call stripe here
	stripe.Key = "stripe_test_key"
	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(int64(payment.Amount * 100)),
		Currency: stripe.String("NZD"),
		Capture:  stripe.Bool(true),
	}
	chargeParams.AddMetadata("payID", payment.ID.Hex())
	chargeParams.SetSource(req.Source)
	chargeParams.SetIdempotencyKey(primitive.NewObjectID().Hex())
	go charge.New(chargeParams)
	result := true
	order, err := db.UpdatePaymentResult(req.PaymentID, result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
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
