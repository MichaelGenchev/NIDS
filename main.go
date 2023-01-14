package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/MichaelGenchev/NIDS/capturer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var mongoURI = "mongodb://localhost:27017"
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	demoDB := client.Database("NIDS")
	err = demoDB.CreateCollection(ctx, "cats")
	if err != nil {
		log.Fatal(err)
	}
	catsCollection := demoDB.Collection("cats")
	result, err := catsCollection.InsertOne(ctx, bson.D{
		{Key:"name", Value: "MOcha"},
		{Key:"type", Value: "Turkish"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	// Run()
}

// func Run() {
// 	capturer.Capture()
// }
