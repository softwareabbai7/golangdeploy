package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/softwareabbai7/test_golangapi/common"
	"github.com/softwareabbai7/test_golangapi/types"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Login struct
	var login types.Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	phoneNumber := login.PhoneNumber
	password := login.Password

	message, err := common.LoginCustomer(phoneNumber, password)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := types.Response{
		Message:    message,
		StatusCode: 200,
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)

}
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Login struct
	var registration types.Registration
	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	message, _ := common.RegisterCustomer(registration)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := types.Response{
		Message:    message,
		StatusCode: 500,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)

}
