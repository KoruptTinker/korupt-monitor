package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/net/context"
)

type KeypressWeeklyAggregationData struct {
	ID         time.Time `bson:"_id"`
	Keypresses int       `bson:"keyPresses"`
}

func (db *Mongo) GetWeeklyKeypressData(ctx context.Context) ([]KeypressWeeklyAggregationData, error) {
	pipeline := buildWeeklyKeypressAggregation()

	res, err := db.KeypressCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var results []KeypressWeeklyAggregationData

	if err := res.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func buildWeeklyKeypressGroupStage(unit string, timezone string) bson.D {
	return bson.D{
		{
			Key: "$group",
			Value: bson.D{
				{
					Key: "_id",
					Value: bson.D{
						{
							Key: "$dateTrunc",
							Value: bson.D{
								{Key: "date", Value: "$timestamp"},
								{Key: "unit", Value: unit},
								{Key: "timezone", Value: timezone},
							},
						},
					},
				},
				{
					Key:   "keyPresses",
					Value: bson.D{{Key: "$sum", Value: "$count"}},
				},
			},
		},
	}
}

func buildWeeklyKeypressAggregation() bson.A {
	var pipeline bson.A

	pipeline = append(pipeline, buildWeeklyMatchStage())
	pipeline = append(pipeline, buildWeeklyKeypressGroupStage("day", "EST"))
	pipeline = append(pipeline, buildSortStage("_id", -1))

	return pipeline
}
