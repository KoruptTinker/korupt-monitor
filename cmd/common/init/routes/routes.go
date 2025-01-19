package routes

import (
	"github.com/KoruptTinker/korupt-monitor/internal/server/controllers"
	"github.com/KoruptTinker/korupt-monitor/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func AddURLs(r *gin.RouterGroup, server controllers.Controller) {
	r.GET("/knockknock", middlewares.RequestHandler(server.HealthCheck))
}
