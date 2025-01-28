package main

import (
	"fmt"
	"time"

	"github.com/KoruptTinker/korupt-monitor/config"
	"github.com/KoruptTinker/korupt-monitor/internal/client/controllers"
	"github.com/go-co-op/gocron/v2"
	hook "github.com/robotn/gohook"
)

func InitClient() gocron.Scheduler {
	engine, err := gocron.NewScheduler()
	if err != nil {
		panic(fmt.Sprintf("Error creating scheduler engine: %v", err.Error()))
	}
	configObj := config.ParseConfig("config/prod.yaml")
	clientController := controllers.New(&configObj)

	go InitRecorders(clientController)

	AddJobs(engine, clientController)

	return engine
}

func AddJobs(engine gocron.Scheduler, client *controllers.ClientController) {
	engine.NewJob(gocron.DurationJob(time.Minute*5), gocron.NewTask(client.TransmitKeypressData))
	engine.NewJob(gocron.DurationJob(time.Minute*5), gocron.NewTask(client.TransmitClickData))
}

func InitRecorders(client *controllers.ClientController) {
	client.RecordClick()
	client.RecordKeyPress()
	s := hook.Start()
	<-hook.Process(s)
}
