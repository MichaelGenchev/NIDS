package main

import (
	"context"
	"log"

	"github.com/MichaelGenchev/NIDS/alert"
	"github.com/MichaelGenchev/NIDS/capturer"
	"github.com/MichaelGenchev/NIDS/parser"
	"github.com/MichaelGenchev/NIDS/sbd"

	alertsRepo "github.com/MichaelGenchev/NIDS/alert/repository"
	parsedPacketsRepo "github.com/MichaelGenchev/NIDS/parser/repository"
	sbdRepo "github.com/MichaelGenchev/NIDS/sbd/repository"

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

	alertsCollection := demoDB.Collection("alerts")
	ppColection := demoDB.Collection("parsedPackets")
	sCollection := demoDB.Collection("signatures")

	AlertsRepo := alertsRepo.NewAlertMongoRepository(alertsCollection)
	PPrepo := parsedPacketsRepo.NewParsedPacketsMongoRepo(ppColection)
	Srepo := sbdRepo.NewSignatureStorage(sCollection)

	chPackets := make(chan gopacket.Packet, 100)
	chPP := make(chan *parser.ParsedPacket, 100)
	chDetection := make(chan sbd.DetectionEvent)

	c := capturer.Capturer{}
	p := parser.NewParser(PPrepo)
	sbd := sbd.NewSBD(Srepo)
	alerter := alert.NewAlerter(AlertsRepo)

	go alerter.ListenForDetectionEvents(chDetection)
	go sbd.AcceptParsedPackets(chPP, chDetection)
	go p.Listen(chPackets, chPP)
	c.Capture(chPackets)

}
