package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Clicks struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp       time.Time          `bson:"timestamp"`
	LeftClickCount  int                `bson:"left_click_count"`
	RightClickCount int                `bson:"right_click_count"`
}

func (db *Mongo) InsertClickData(ctx context.Context, leftCount int, rightCount int) error {
	timeNow := time.Now()
	clickData := Clicks{
		Timestamp:       timeNow,
		LeftClickCount:  leftCount,
		RightClickCount: rightCount,
	}

	_, err := db.ClickCollection.InsertOne(ctx, clickData)

	return err
}
