package common

import "github.com/KoruptTinker/korupt-monitor/internal/client/controllers"

func InitCron() *controllers.ClientController {
	client := controllers.New()
	addSchedulerJobs(client)
	return client
}

func addSchedulerJobs(client *controllers.ClientController) {
}
