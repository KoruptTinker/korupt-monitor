package v1

import (
	"github.com/KoruptTinker/korupt-monitor/internal/server/controllers"
	"github.com/KoruptTinker/korupt-monitor/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

func AddURLs(r *gin.RouterGroup, server *controllers.Controller) {
	r.PUT("/keypresses", middlewares.RequestHandler(server.RecordKeyPressData))
	r.PUT("/clicks", middlewares.RequestHandler(server.RecordUserClick))
	r.GET("/weekly", middlewares.RequestHandler(server.DisplayWeeklyData))
}
