package main

import (
	"fmt"
	"time"

	"github.com/KoruptTinker/korupt-monitor/internal/client/controllers"
	"github.com/go-co-op/gocron/v2"
)

func InitClient() gocron.Scheduler {
	engine, err := gocron.NewScheduler()
	if err != nil {
		panic(fmt.Sprintf("Error creating scheduler engine: %v", err.Error()))
	}
	clientController := controllers.New()

	AddJobs(engine, clientController)

	return engine
}

func AddJobs(engine gocron.Scheduler, client *controllers.ClientController) {
	engine.NewJob(gocron.DurationJob(time.Duration(time.Minute*5)), gocron.NewTask(client.TransmitClickData()))
	engine.NewJob(gocron.DurationJob(time.Duration(time.Minute*2)), gocron.NewTask(client.TransmitKeypressData()))
}
