package common

import (
	"github.com/KoruptTinker/korupt-monitor/cmd/common/init/routes"
	"github.com/KoruptTinker/korupt-monitor/config"
	"github.com/KoruptTinker/korupt-monitor/internal/server/controllers"
	"github.com/gin-gonic/gin"
)

func InitServer() (*gin.Engine, *config.Config) {
	services := initServices()
	configuration := config.ParseConfig("config/prod.yaml")
	server := controllers.Controller{
		Services: services,
	}
	engine := gin.Default()
	routes.AddURLs(&engine.RouterGroup, server)

	return engine, &configuration
}

func initServices() controllers.External {
	return controllers.External{}
}
