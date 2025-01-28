package controllers

import (
	"fmt"
	"sync"

	"github.com/KoruptTinker/korupt-monitor/config"
	"github.com/KoruptTinker/korupt-monitor/internal/services"
	korupt_monitor_server "github.com/KoruptTinker/korupt-monitor/internal/services/korupt-monitor-server"
	"github.com/go-co-op/gocron/v2"
)

type ClientController struct {
	Engine         gocron.Scheduler
	UserClicks     ClickData
	UserKeyPresses KeyboardData
	External       services.External
}

func New(config *config.Config) *ClientController {
	sched, err := gocron.NewScheduler()
	if err != nil {
		panic(fmt.Sprintf("Error creating cron engine: %v", err.Error()))
	}
	return &ClientController{
		Engine: sched,
		UserKeyPresses: KeyboardData{
			KeyPresses: 0,
			Lock:       &sync.Mutex{},
		},
		UserClicks: ClickData{
			LeftClicks:  0,
			RightClicks: 0,
			Lock:        &sync.Mutex{},
		},
		External: services.External{
			KoruptMonitorServer: *korupt_monitor_server.New(config),
		},
	}
}
