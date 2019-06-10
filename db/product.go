package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

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

//ReadProducts func
func ReadProducts(ids []primitive.ObjectID) ([]*Product, error) {
	var products []*Product
	query := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := DB.Collection("products").Find(CTX, query)
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

//ReadAllProducts func
func ReadAllProducts() ([]*Product, error) {
	var products []*Product
	cursor, err := DB.Collection("products").Find(CTX, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(CTX) {
		p := &Product{}
		if err := cursor.Decode(p); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
