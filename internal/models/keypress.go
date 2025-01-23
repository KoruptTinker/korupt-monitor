package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyPresses struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp"`
	Count     int                `bson:"count"`
}

func (db *Mongo) InsertKeypressData(ctx context.Context, count int) error {
	timeNow := time.Now()
	keyData := KeyPresses{
		Timestamp: timeNow,
		Count:     count,
	}

	_, err := db.KeypressCollection.InsertOne(ctx, keyData)

	return err
}
