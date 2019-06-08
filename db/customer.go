package db

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
)

// var collection *mongo.Collection

// func init() {
// 	collection := Client.Database("eshop").Collection("customers")
// }

//AddCustomer func
func AddCustomer(customer *Customer) error {
	if _, err := readCustomerByEmail(customer.Email); err == nil {
		return errors.New("user already exists")
	}
	_, err := DB.Collection("customers").InsertOne(CTX, customer)
	if err != nil {
		return err
	}
	return nil
}

//Authenticate func
func Authenticate(email, password string) (*Customer, error) {
	customer, err := readCustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		return nil, err
	}
	return customer, nil
}

//check customer email occupied
func readCustomerByEmail(email string) (*Customer, error) {
	customer := &Customer{}
	query := bson.M{"email": email}
	if err := DB.Collection("customers").FindOne(CTX, query).Decode(customer); err != nil {
		return nil, err
	}
	return customer, nil
}
