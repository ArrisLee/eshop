package db

//AddCustomer func
func AddCart(cart *Cart) error {
	_, err := DB.Collection("carts").InsertOne(CTX, cart)
	if err != nil {
		return err
	}
	return nil
}
