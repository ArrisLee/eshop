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
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Desc  string             `json:"desc" bson:"desc"`
	Img   string             `json:"img" bson:"img"`
	Price float64            `json:"price" bson:"price"`
}

//Customer struct
type Customer struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
	Token    string             `json:"token" bson:"token"`
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
