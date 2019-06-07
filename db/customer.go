package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//AddCustomer func
func AddCustomer(customer *Customer) error {
	//to-do need to set email as unique index
	collection := Client.Database("eshop").Collection("customers")
	_, err := collection.InsertOne(CTX, customer)
	if err != nil {
		return err
	}
	return nil
}

//Authenticate func
func Authenticate(email string, password string) (*Customer, error) {
	log.Println(email, password)
	customer := &Customer{}
	collection := Client.Database("eshop").Collection("customers")
	query := bson.M{"email": email, "password": password}
	if err := collection.FindOne(CTX, query).Decode(customer); err != nil {
		return nil, err
	}
	return customer, nil
}
