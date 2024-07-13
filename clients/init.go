package clients

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

func InitFireStoreClient() *firestore.Client {

	if firestoreClient != nil {
		return firestoreClient
	}
	encodedString := os.Getenv("casual_file")

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)

	}

	// Convert bytes to string
	jsonFile := []byte(string(decodedBytes))
	opt := option.WithCredentialsJSON(jsonFile)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	firestoreClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v", err)
	}
	return firestoreClient

}

func CloseFirestoreClient() {
	if firestoreClient != nil {
		err := firestoreClient.Close()
		if err != nil {
			log.Printf("Failed to close Firestore client: %v", err)
		}
	}
}
