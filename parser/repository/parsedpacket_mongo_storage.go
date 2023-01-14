package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/MichaelGenchev/NIDS/parser"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type ParsedPacketMongoRepository struct {
	collection *mongo.Collection
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
func (r *ParsedPacketMongoRepository) FindByID(id string) (*parser.ParsedPacket, error) {
    var parsedPacket parser.ParsedPacket

	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	defer cancel()

    filter := bson.M{"ID": id}

    err := r.collection.FindOne(ctx, filter).Decode(&parsedPacket)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("parsed packet with ID %s not found", id)
	}
    if err != nil {
        return nil, fmt.Errorf("parsed packet repo find err: %w", err)
    }

    return &parsedPacket, nil
}

func (r *ParsedPacketMongoRepository) FindAll() ([]*parser.ParsedPacket, error) {
    var parsedPackets []*parser.ParsedPacket

    ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
    defer cancel()

    cur, err := r.collection.Find(ctx, bson.D{})
    if err != nil {
        return nil, fmt.Errorf("parsed packet repo err: %w", err)
    }
    defer cur.Close(ctx)

    for cur.Next(ctx) {
        var parsedPacket parser.ParsedPacket
        if err := cur.Decode(&parsedPacket); err != nil {
            return nil, fmt.Errorf("parsed packet repo err: %w", err)
        }
        parsedPackets = append(parsedPackets, &parsedPacket)
    }
    if err := cur.Err(); err != nil {
        return nil, fmt.Errorf("parsed packet repo err: %w", err)
    }

    return parsedPackets, nil
}

func (r *ParsedPacketMongoRepository) DeleteByID(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
    defer cancel()

    filter := bson.M{"ID": id}

    res, err := r.collection.DeleteOne(ctx, filter)
    if err != nil {
        return fmt.Errorf("parsed packet repo err: %w", err)
    }
    if res.DeletedCount == 0 {
        return fmt.Errorf("parsed packet with ID %s not found", id)
    }
    return nil
}

func (r *ParsedPacketMongoRepository) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, fmt.Errorf("parsed packet repo err: %w", err)
	}

	return int(count), nil
}