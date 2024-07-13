package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/softwareabbai7/test_golangapi/common"
	"github.com/softwareabbai7/test_golangapi/types"
)

func AddEditProductHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Login struct
	var product_details types.ProductListing
	var message string
	err := json.NewDecoder(r.Body).Decode(&product_details)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if product_details.ProductCollectionID == "" {
		message, err = common.AddProductHelper(product_details)
	} else {
		message, err = common.EditProducts(product_details)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := types.Response{
		Message:    message,
		StatusCode: 200,
	}
	if err != nil {
		resp = types.Response{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)

}
