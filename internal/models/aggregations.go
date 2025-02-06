package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func buildWeeklyMatchStage() bson.D {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	startDate := today.AddDate(0, 0, -6)
	return bson.D{
		{
			Key: "$match", Value: bson.D{
				{
					Key: "timestamp", Value: bson.D{{Key: "$gte", Value: startDate}},
				},
			},
		},
	}
}

func buildSortStage(field string, order int) bson.D {
	return bson.D{
		{
			Key: "$sort",
			Value: bson.D{{
				Key:   field,
				Value: order,
			}},
		},
	}
}
