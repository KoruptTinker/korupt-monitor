package models

import (
	"fmt"

	"github.com/KoruptTinker/korupt-monitor/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client              *mongo.Client
	DB                  *mongo.Database
	ClickCollection     *mongo.Collection
	KeypressCollection  *mongo.Collection
	TImeshareCollection *mongo.Collection
}

func NewMongoClient(config config.Config) *Mongo {
	mongoConfig := config.Db
	connectionURL := fmt.Sprintf("mongodb+srv://%s:%s@%s", mongoConfig.User, mongoConfig.Pass, mongoConfig.ConectionURL)

	client, err := mongo.Connect(options.Client().ApplyURI(connectionURL))
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB. Error: %v", err.Error()))
	}

	dbHandle := client.Database("korupt-monitor")
	clickHandle := dbHandle.Collection("clicks")
	keypressHandle := dbHandle.Collection("keypresses")
	timeshareHandle := dbHandle.Collection("timeshrae")

	mongoObj := Mongo{
		Client:              client,
		DB:                  dbHandle,
		ClickCollection:     clickHandle,
		KeypressCollection:  keypressHandle,
		TImeshareCollection: timeshareHandle,
	}

	return &mongoObj
}
