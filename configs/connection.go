package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
)

func ConnectDB() *mongo.Client {
	// mongodb new instance
	client, err := mongo.NewClient(options.Client().ApplyURI(LoadENV("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}
	
	// timeout configuration
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err = client.Connect(ctx); err != nil {
			log.Fatal(err)
		}
		
		// defer client.Disconnect(ctx)
		
		// ping the database
		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			log.Fatal(err)
		}

		fmt.Println("===> MONGODB Connected successfully!! <===")

			// printall collection in DB
		databases, err := client.ListDatabaseNames(ctx, bson.M{})

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(databases)

		
		return client 
		
}

var DB *mongo.Client = ConnectDB()

// get all database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golang-mongo").Collection(collectionName)

	return collection
}