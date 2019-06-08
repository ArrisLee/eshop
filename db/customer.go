package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"

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
func Authenticate(email, password string) (*Customer, error) {
	log.Println(email, password)
	customer := &Customer{}
	collection := Client.Database("eshop").Collection("customers")
	query := bson.M{"email": email}
	if err := collection.FindOne(CTX, query).Decode(customer); err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		return nil, err
	}
	return customer, nil
}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
