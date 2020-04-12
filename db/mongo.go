package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/test?retryWrites=true&w=majority",
	))

	if err != nil {
		log.Fatal(err)
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("New MongoDB client!", client.NumberSessionsInProgress())

	return nil
}

func GetClient() *mongo.Client {
	return client
}
