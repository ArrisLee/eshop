package controller

import (
	"eshop/db"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cartRequest struct {
	products []*db.Product `json:"products"`
}

//CreateCart func
func CreateCart(c echo.Context) error {
	req := &cartRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	productMap := make(map[primitive.ObjectID]int32)
	var ids []primitive.ObjectID
	for _, p := range req.products {
		productMap[p.ID] = p.Quantity
		ids = append(ids, p.ID)
	}
	products, err := db.ReadProducts(ids)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var totalPrice float64
	for _, p := range products {
		totalPrice += p.Price * float64(productMap[p.ID])
	}
	cart := db.Cart{}
	cart.ID = primitive.NewObjectID()
	cart.Products = req.products
	cart.TotalPrice = totalPrice
	log.Println(cart)
	return c.JSON(http.StatusCreated, cart.ID)
}
