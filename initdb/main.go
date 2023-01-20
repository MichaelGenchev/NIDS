package main

import (
	"context"
	"flag"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = flag.String("uri", "mongodb://localhost:27017", "MongoDB URI")

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	// Create collections
	collections := []string{"alerts", "parsedPackets", "signatures"}
	for _, collectionName := range collections {
		client.Database("NIDS").CreateCollection(context.TODO(), collectionName)
	}
}
