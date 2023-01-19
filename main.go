package main

import (
	"log"

	"github.com/MichaelGenchev/NIDS/cli"
	"github.com/MichaelGenchev/NIDS/config"
	"github.com/MichaelGenchev/NIDS/setup"
)

func main() {
	cli := cli.ParseFlags()

	cfg, err := config.LoadConfig(cli.InterfaceFlag, cli.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	setup.Start(cfg)
}
