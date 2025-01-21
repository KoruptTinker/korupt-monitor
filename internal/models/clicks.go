package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Clicks struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp"`
	Count     int                `bson:"count"`
}

func (db *Mongo) InsertClickData(ctx context.Context, count int) error {
	timeNow := time.Now()
	clickData := Clicks{
		Timestamp: timeNow,
		Count:     count,
	}

	_, err := db.ClickCollection.InsertOne(ctx, clickData)

	return err
}
