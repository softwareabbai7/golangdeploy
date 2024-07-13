package common

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/softwareabbai7/test_golangapi/clients"
	"github.com/softwareabbai7/test_golangapi/types"
)

func getLoginDetails(phone_number string) (*firestore.DocumentSnapshot, error) {
	firestoreClient := clients.InitFireStoreClient()
	data, err := firestoreClient.Collection(FireStoreCustomerMain).Doc(phone_number).Get(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil

}

func LoginCustomer(phone_number, password string) (string, error) {

	data, err := getLoginDetails(phone_number)
	if err != nil {
		return "user_not_found", err
	}
	data_password := data.Data()["Password"]
	if data_password == password {
		return "success", nil
	}
	return "wrong_password", nil

}
func RegisterCustomer(registration types.Registration) (string, error) {
	firestoreClient := clients.InitFireStoreClient()
	_, err := getLoginDetails(registration.PhoneNumber)
	if err != nil {
		_, err = firestoreClient.Collection(FireStoreCustomerMain).Doc(registration.PhoneNumber).Set(context.Background(), registration)
		if err != nil {
			return "reg_failed", err
		}

		return "reg_success", nil
	}
	return "user_exist", nil

}
