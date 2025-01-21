package api

import (
	v1 "github.com/KoruptTinker/korupt-monitor/cmd/common/init/routes/api/v1"
	"github.com/KoruptTinker/korupt-monitor/internal/server/controllers"
	"github.com/gin-gonic/gin"
)

func AddURLs(r *gin.RouterGroup, server *controllers.Controller) {
	v1Router := r.Group("/v1")
	v1.AddURLs(v1Router, server)
}
