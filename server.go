package main

import (
	"context"
	"eshop/controller"
	"eshop/db"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	routers := gin.Default()
	//---------- connect db -------------------
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://lindsay:6666666666@arris-wn60t.mongodb.net/test"))
	log.Println(client.Ping(ctx, readpref.Primary()))
	if err != nil {
		log.Fatal("mongodb connection failed")
	}
	//---------- populate product struct ------------------
	p := db.Product{}
	p.ID = primitive.NewObjectID()
	p.Name = "haha"
	p.Pic = "lala"
	p.Price = 111.1

	//---------- insert product to DB
	client.Database("eshop").Collection("products").InsertOne(ctx, p)
	// c := db.Customer{}
	// c.Name = "lindsay"
	// client.Database("eshop").Collection("cutomers").FindOne(ctx, c)
	// Databse = client.Database("Eshop")
	routers.GET("/api/v1/products", controller.GetAllProducts)
	routers.GET("/api/v1/customers", controller.GetAllCustomers)
	routers.POST("/api/v1/order", controller.CreateOrder)
	routers.Run(":8000")
}
