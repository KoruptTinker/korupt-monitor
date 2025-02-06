package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/net/context"
)

type ClickWeeklyAggregationData struct {
	ID              time.Time `bson:"_id"`
	LeftClickCount  int       `bson:"leftClicks"`
	RightClickCount int       `bson:"rightClicks"`
}

func (db *Mongo) GetWeeklyClickData(ctx context.Context) ([]ClickWeeklyAggregationData, error) {
	pipeline := buildWeeklyClicksAggregation()

	res, err := db.ClickCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var results []ClickWeeklyAggregationData

	if err := res.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func buildWeeklyClickGroupStage(unit string, timezone string) bson.D {
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
					Key:   "leftClicks",
					Value: bson.D{{Key: "$sum", Value: "$left_click_count"}},
				},
				{
					Key:   "rightClicks",
					Value: bson.D{{Key: "$sum", Value: "$right_click_count"}},
				},
			},
		},
	}
}

func buildWeeklyClicksAggregation() bson.A {
	var pipeline bson.A

	pipeline = append(pipeline, buildWeeklyMatchStage())
	pipeline = append(pipeline, buildWeeklyClickGroupStage("day", "EST"))
	pipeline = append(pipeline, buildSortStage("_id", -1))

	return pipeline
}
