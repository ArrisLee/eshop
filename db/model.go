package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb+srv://arris:871102@arris-wn60t.mongodb.net/test"

var (
	//CTX shared
	CTX context.Context
	//DB shared
	DB *mongo.Database
)

//Product struct
type Product struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Desc     string             `json:"desc" bson:"desc"`
	Img      string             `json:"img" bson:"img"`
	Price    float64            `json:"price" bson:"price"`
	Quantity int32              `json:"quantity" bson:"quantity"`
}

//Customer struct
type Customer struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Token    string             `json:"token" bson:"token"`
}

//Cart struct
type Cart struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID primitive.ObjectID `json:"customerID" bson:"customerID"`
	Products   []*Product         `json:"products" bson:"products"`
	TotalPrice float64            `json:"totalPrice" bson:"totalPrice"`
}

//Order struct
type Order struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID primitive.ObjectID `json:"customerID" bson:"customerID"`
	PaymentID  primitive.ObjectID `json:"paymentID" bson:"paymentID"`
	ShortID    string             `json:"shortID" bson:"shortID"`
	Name       string             `json:"name" bson:"name"`
	Address    string             `json:"address" bson:"address"`
	Phone      string             `json:"phone" bson:"phone"`
	Cart       *Cart              `json:"cart" bson:"cart"`
	Price      float64            `json:"price" bson:"price"`
	Paid       bool               `json:"paid" bson:"paid"`
	Dispatched bool               `json:"dispatched" bson:"dispatched"`
	Delivered  bool               `json:"delivered" bson:"delivered"`
}

//Payment struct
type Payment struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	OrderID    primitive.ObjectID `json:"orderID" bson:"orderID"`
	CustomerID primitive.ObjectID `json:"customerID" bson:"customerID"`
	Type       string             `json:"type" bson:"type"`
	Amount     float64            `json:"amount" bson:"amount"`
	Success    bool               `json:"success" bson:"success"`
}

func init() {
	CTX, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(CTX, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = client.Database("eshop")
}
