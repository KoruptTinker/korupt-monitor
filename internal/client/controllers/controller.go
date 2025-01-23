package controllers

import (
	"fmt"

	"github.com/KoruptTinker/korupt-monitor/internal/services"
	"github.com/go-co-op/gocron/v2"
)

type ClientController struct {
	Engine         gocron.Scheduler
	UserClicks     ClickData
	UserKeyPresses KeyboardData
	External       services.External
}

func New() *ClientController {
	sched, err := gocron.NewScheduler()
	if err != nil {
		panic(fmt.Sprintf("Error creating cron engine: %v", err.Error()))
	}
	return &ClientController{
		Engine:         sched,
		UserKeyPresses: KeyboardData{},
		UserClicks:     ClickData{},
		External:       services.External{},
	}
}
