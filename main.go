package main

import (
	"log"

	"github.com/MichaelGenchev/NIDS/cli"
	"github.com/MichaelGenchev/NIDS/config"
	"github.com/MichaelGenchev/NIDS/setup"
)

// Make CLI channel to send information to the CLI interface
func main() {
	cli := cli.ParseFlags()
	go cli.AcceptInfo()
	cfg, err := config.LoadConfig(cli.InterfaceFlag, cli.MongoURI, cli.Channel)
	if err != nil {
		log.Fatal(err)
	}
	setup.Start(cfg)
}
