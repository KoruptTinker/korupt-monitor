package middlewares

import (
	"github.com/KoruptTinker/korupt-monitor/internal/server/controllers"
	"github.com/gin-gonic/gin"
)

func RequestHandler(controllerFunc func(c *gin.Context) controllers.BaseResponse) func(c *gin.Context) {
	return func(c *gin.Context) {
		response := controllerFunc(c)
		c.JSON(response.ResponseCode, response.ResponseData)
	}
}
