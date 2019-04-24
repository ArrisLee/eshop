package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var (
// 	Databse *mongo.Database
// 	Ctx     context.Context
// )

//Product struct
type Product struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Pic   string             `json:"pic" bson:"pic"`
	Price float32            `json:"price" bson:"price"`
}

type Customer struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Phone string             `json:"phone" bson:"phone"`
}

// func init() {
// 	Ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
// 	client, err := mongo.Connect(Ctx, options.Client().ApplyURI("mongodb+srv://lindsay:6666666666@arris-wn60t.mongodb.net/test"))
// 	log.Println(client.Ping(Ctx, readpref.Primary()))
// 	if err != nil {
// 		log.Fatal("mongodb connection failed")
// 	}
// 	Databse = client.Database("Eshop")
// }
