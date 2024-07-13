package common

import (
	"context"

	"github.com/softwareabbai7/test_golangapi/clients"
	"github.com/softwareabbai7/test_golangapi/types"
)

func AddProductHelper(product_details types.ProductListing) (string, error) {

	firestoreClient := clients.InitFireStoreClient()
	docref, _, err := firestoreClient.Collection(FireStroeProductWareHouseOne).Doc(FireStoreProductMain).Collection(product_details.ProductCategory).Add(context.Background(), product_details)

	if err != nil {
		return "", err
	}
	product_details.ProductCollectionID = docref.ID
	docref.Set(context.Background(), product_details)
	return "product_inserted", nil
}
func EditProducts(product_details types.ProductListing) (string, error) {
	firestoreClient := clients.InitFireStoreClient()
	_, err := firestoreClient.Collection(FireStroeProductWareHouseOne).Doc(FireStoreProductMain).Collection(product_details.ProductCategory).Doc(product_details.ProductCollectionID).Set(context.Background(), product_details)
	if err != nil {
		return "product_not_exist", err
	}
	return "updated_product", nil

}
