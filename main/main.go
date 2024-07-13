package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/softwareabbai7/test_golangapi/clients"
	"github.com/softwareabbai7/test_golangapi/handlers"
)

func main() {

	clients.InitFireStoreClient()
	router := mux.NewRouter()
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegistrationHandler).Methods("POST")
	//Admin function

	router.HandleFunc("/addproduct", handlers.AddEditProductHandler).Methods("POST")
	router.HandleFunc("/updateproduct", handlers.AddEditProductHandler).Methods("POST")
	router.HandleFunc("/ping", PingHandler).Methods("GET")
	fmt.Println("started")
	http.ListenAndServe(":8080", router)

}
func PingHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"response": "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
