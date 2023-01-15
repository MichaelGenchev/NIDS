package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/MichaelGenchev/NIDS/sbd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignatureStorage struct {
	collection *mongo.Collection
}

func NewSignatureStorage(collection *mongo.Collection) *SignatureStorage {
	return &SignatureStorage{collection: collection}
}

func (s *SignatureStorage) FindAll() ([]*sbd.Signature, error) {
	var signatures []*sbd.Signature

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("parsed packet repo err: %w", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var signature sbd.Signature
		if err := cur.Decode(&signature); err != nil {
			return nil, fmt.Errorf("parsed packet repo err: %w", err)
		}
		signatures = append(signatures, &signature)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("parsed packet repo err: %w", err)
	}

	return signatures, nil
}
