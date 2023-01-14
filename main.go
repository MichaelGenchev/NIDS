package main

import (
	"context"
	"log"

	"github.com/MichaelGenchev/NIDS/capturer"
	"github.com/MichaelGenchev/NIDS/parser"
	parsedPacketsRepo "github.com/MichaelGenchev/NIDS/parser/repository"
	"github.com/google/gopacket"
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
	err = demoDB.CreateCollection(ctx, "parsedPackets")
	if err != nil {
		log.Fatal(err)
	}
	ppColection := demoDB.Collection("parsedPackets")
	
	repo := parsedPacketsRepo.NewParsedPacketsMongoRepo(ppColection)

	chPackets := make(chan gopacket.Packet, 100)
	c := capturer.Capturer{}
	p := parser.NewParser(repo)
	go p.Listen(chPackets)
	c.Capture(chPackets)
	
}


