package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddPayment func
func AddPayment(payment *Payment) error {
	_, err := DB.Collection("payments").InsertOne(CTX, payment)
	if err != nil {
		return err
	}
	return nil
}

//UpdatePaymentResult func
func UpdatePaymentResult(paymentID primitive.ObjectID, result bool) (*Order, error) {
	query := bson.M{"_id": paymentID}
	update := bson.M{"$set": bson.M{"success": result}}
	log.Println(query, update)
	_, err := DB.Collection("payments").UpdateOne(CTX, query, update)
	if err != nil {
		return nil, err
	}
	order, err := UpdateOrderPayment(paymentID, result)
	if err != nil {
		return nil, err
	}
	return order, nil
}

//ReadPaymentsByCustomerID func
func ReadPaymentsByCustomerID(customerID primitive.ObjectID) ([]*Payment, error) {
	var payments []*Payment
	query := bson.M{"customerID": customerID}
	cursor, err := DB.Collection("payments").Find(CTX, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(CTX) {
		p := &Payment{}
		if err := cursor.Decode(p); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}
