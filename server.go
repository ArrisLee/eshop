package main

import (
	"context"
	"eshop/controller"
	"eshop/db"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	routers := gin.Default()
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://lindsay:6666666666@arris-wn60t.mongodb.net/test"))
	log.Println(client.Ping(ctx, readpref.Primary()))
	if err != nil {
		log.Fatal("mongodb connection failed")
	}
	p := db.Product{}
	p.ID = primitive.NewObjectID()
	p.Name = "haha"
	p.Pic = "lala"
	p.Price = 111.1
	client.Database("test").Collection("check").InsertOne(ctx, p)
	// Databse = client.Database("Eshop")
	routers.GET("/api/v1/products", controller.GetAllProducts)
	routers.Run(":8000")
}
