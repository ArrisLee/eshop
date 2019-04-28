package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb+srv://lindsay:6666666666@arris-wn60t.mongodb.net/test"

var (
	CTX    context.Context
	Client *mongo.Client
)

//Product struct
type Product struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Pic   string             `json:"pic" bson:"pic"`
	Price float32            `json:"price" bson:"price"`
}

//Customer struct
type Customer struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Phone string             `json:"phone" bson:"phone"`
}

func init() {
	CTX, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	Client, err = mongo.Connect(CTX, options.Client().ApplyURI("mongodb+srv://lindsay:6666666666@arris-wn60t.mongodb.net/test"))
	if err != nil {
		log.Fatal("data base connection failed")
	}
}
