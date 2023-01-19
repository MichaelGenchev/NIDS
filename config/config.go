package config

import (
	"errors"
	"github.com/MichaelGenchev/NIDS/cli"
)

var (
	ErrEmptyNetworkInterface = errors.New("network interface is empty")
	ErrEmptyMongoURI         = errors.New("mongo uri is empty")
)

type Config struct {
	NetworkInterface string
	MongoURI         string
	CLIChannel       chan cli.Info
}

func LoadConfig(networkInterface, mongoURI string, cliCh chan cli.Info) (*Config, error) {
	if networkInterface == "" {
		return nil, ErrEmptyNetworkInterface
	}
	if mongoURI == "" {
		return nil, ErrEmptyMongoURI
	}
	return &Config{
		NetworkInterface: networkInterface,
		MongoURI:         mongoURI,
		CLIChannel:       cliCh,
	}, nil
}
