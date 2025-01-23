package common

import (
	"github.com/KoruptTinker/korupt-monitor/cmd/common/init/routes"
	"github.com/KoruptTinker/korupt-monitor/config"
	httpClient "github.com/KoruptTinker/korupt-monitor/internal/core/http_client"
	"github.com/KoruptTinker/korupt-monitor/internal/models"
	"github.com/KoruptTinker/korupt-monitor/internal/server/controllers"
	"github.com/KoruptTinker/korupt-monitor/internal/services"
	korupt_monitor_server "github.com/KoruptTinker/korupt-monitor/internal/services/korupt-monitor-server"
	"github.com/gin-gonic/gin"
)

func InitServer() (*gin.Engine, *config.Config) {
	configuration := config.ParseConfig("config/prod.yaml")
	services := initServices(configuration)
	mongo := models.NewMongoClient(configuration)
	server := controllers.Controller{
		DB:       *mongo,
		Services: services,
	}
	engine := gin.Default()
	routes.AddURLs(&engine.RouterGroup, &server)

	return engine, &configuration
}

func initServices(config config.Config) services.External {
	return services.External{
		KoruptMonitorServer: korupt_monitor_server.Service{
			BaseExternal: httpClient.BaseExternal{
				Hostname: config.External.KoruptMonitorServer.Hostname,
			},
		},
	}
}
