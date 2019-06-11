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
func UpdatePaymentResult(paymentID primitive.ObjectID, result bool) error {
	query := bson.M{"_id": paymentID}
	update := bson.M{"$set": bson.M{"success": result}}
	log.Println(query, update)
	_, err := DB.Collection("payments").UpdateOne(CTX, query, update)
	if err != nil {
		return err
	}
	if err := UpdateOrderPayment(paymentID, result); err != nil {
		return err
	}
	return nil
}
