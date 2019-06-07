package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//AddProduct func
func AddProduct(product *Product) (*Product, error) {
	collection := Client.Database("eshop").Collection("products")
	_, err := collection.InsertOne(CTX, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

//ReadAllProducts func
func ReadAllProducts() ([]*Product, error) {
	var products []*Product
	collection := Client.Database("eshop").Collection("products")
	cursor, err := collection.Find(CTX, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(CTX) {
		p := &Product{}
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	return products, nil
}
