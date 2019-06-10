package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddOrder func
func AddOrder(order *Order) error {
	_, err := DB.Collection("orders").InsertOne(CTX, order)
	if err != nil {
		return err
	}
	return nil
}

//ReadOrdersByCustomerID func
func ReadOrdersByCustomerID(customerID primitive.ObjectID) ([]*Order, error) {
	var orders []*Order
	query := bson.M{"customerID": customerID}
	cursor, err := DB.Collection("orders").Find(CTX, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(CTX) {
		o := &Order{}
		if err := cursor.Decode(o); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
