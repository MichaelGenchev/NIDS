package repository

import (
	"context"
	"time"

	"github.com/MichaelGenchev/NIDS/parser"

	"go.mongodb.org/mongo-driver/mongo"
)

type ParsedPacketMongoRepository struct {
	client mongo.Client
	collection mongo.Collection

}


func (r *ParsedPacketMongoRepository) Save(parsedPacket parser.ParsedPacket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, parsedPacket)
	if err != nil {
		return err
	}
	return nil
}