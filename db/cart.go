package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddCart func
func AddCart(cart *Cart) error {
	_, err := DB.Collection("carts").InsertOne(CTX, cart)
	if err != nil {
		return err
	}
	return nil
}

//ReadCart func
func ReadCart(id primitive.ObjectID) (*Cart, error) {
	cart := &Cart{}
	query := bson.M{"_id": id}
	if err := DB.Collection("carts").FindOne(CTX, query).Decode(cart); err != nil {
		return nil, err
	}
	return cart, nil
}
