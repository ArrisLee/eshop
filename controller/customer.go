package controller

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"

	"github.com/ArrisLee/Eshop/db"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type customerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Register func
func Register(c echo.Context) error {
	req := &customerRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, errors.New("invalid parameters"))
	}
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	hasher := md5.New()
	hasher.Write(hashedBytes)
	token := hex.EncodeToString(hasher.Sum(nil))
	customer := &db.Customer{}
	customer.ID = primitive.NewObjectID()
	customer.Email = strings.ToLower(req.Email)
	customer.Password = string(hashedBytes)
	customer.Name = req.Name
	customer.Token = token
	if err := db.AddCustomer(customer); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "success")
}

//Login func
func Login(c echo.Context) error {
	req := &customerRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, errors.New("invalid parameters"))
	}
	customer, err := db.Authenticate(strings.ToLower(req.Email), req.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	customer.Password = ""
	return c.JSON(http.StatusOK, customer)
}

//Authorize func
func Authorize(c echo.Context) bool {
	id := c.Request().Header.Get("id")
	token := c.Request().Header.Get("token")
	customerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	return db.Authorize(customerID, token)
}
