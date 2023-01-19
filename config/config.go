package config

import "errors"

var (
	ErrEmptyNetworkInterface = errors.New("network interface is empty")
	ErrEmptyMongoURI = errors.New("mongo uri is empty")

)
type Config struct {
	NetworkInterface string 
	MongoURI   string 
}

func LoadConfig(networkInterface, mongoURI string) (*Config,error){
	if networkInterface == ""{
		return nil, ErrEmptyNetworkInterface
	}
	if mongoURI == ""{
		return nil, ErrEmptyMongoURI
	}
	return &Config{
		NetworkInterface: networkInterface,
		MongoURI: mongoURI,
	}, nil
}