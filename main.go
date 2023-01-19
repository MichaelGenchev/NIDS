package main

import (
	"log"

	"github.com/MichaelGenchev/NIDS/cli"
	"github.com/MichaelGenchev/NIDS/config"
	"github.com/MichaelGenchev/NIDS/setup"
)


var mongoURI = "mongodb://localhost:27017"
func main() {
	cli.ParseFlags()

	if *cli.Verbose{
		cli.PrintUsage()
		return 
	}
	cfg, err := config.LoadConfig("en0", mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	setup.Start(cfg)
}
