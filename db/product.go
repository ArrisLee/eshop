package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//AddProduct func
func AddProduct(product *Product) (*Product, error) {
	_, err := DB.Collection("products").InsertOne(CTX, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

//ReadAllProducts func
func ReadAllProducts() ([]*Product, error) {
	var products []*Product
	cursor, err := DB.Collection("products").Find(CTX, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(CTX) {
		p := &Product{}
		if err := cursor.Decode(&p); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	return products, nil
}
