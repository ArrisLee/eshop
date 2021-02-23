package controller

import (
	"net/http"
	"net/smtp"

	"github.com/labstack/echo"
)

//SendEmailNotification func, using smtp send
func SendEmailNotification(c echo.Context) error {
	name := c.QueryParam("name")
	email := c.QueryParam("email")
	orderNumber := c.QueryParam("order")
	body := "Hi " + name + ",\n\nHere is your new order: " + orderNumber + ".\nYour goods will be dispatched within 2 working days, we hope you will enjoy it!\n\nBest regards\n\nFresh Team"
	from := "sample@gmail.com"
	pass := "password"
	to := email
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: Greetings from Fresh\n" + body
	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}
