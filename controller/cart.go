package controller

import (
	"fmt"
	"net/http"

	"github.com/ArrisLee/Eshop/db"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cartRequest struct {
	Products []*db.Product `json:"products"`
}

//CreateCart func
func CreateCart(c echo.Context) error {
	if !Authorize(c) {
		return c.JSON(http.StatusUnauthorized, "forbidden")
	}

	req := &cartRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	productMap := make(map[primitive.ObjectID]int32)
	var ids []primitive.ObjectID
	for _, p := range req.Products {
		productMap[p.ID] = p.Quantity
		ids = append(ids, p.ID)
	}
	products, err := db.ReadProducts(ids)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var totalPrice float64
	for _, p := range products {
		p.Desc = ""
		p.Quantity = productMap[p.ID]
		totalPrice += p.Price * float64(productMap[p.ID])
	}
	id := c.Request().Header.Get("id")
	cart := &db.Cart{}
	cart.CustomerID, _ = primitive.ObjectIDFromHex(id)
	cart.ID = primitive.NewObjectID()
	cart.Products = products
	fs := fmt.Sprintf("%.2f", totalPrice)
	cart.TotalPrice, _ = valid.ToFloat(fs)
	if err := db.AddCart(cart); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, cart.ID)
}
