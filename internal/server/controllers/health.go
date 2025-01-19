package controllers

import "github.com/gin-gonic/gin"

func (c *Controller) HealthCheck(ctx *gin.Context) BaseResponse {
	return BaseResponse{
		ResponseCode: 200,
		ResponseData: map[string]interface{}{
			"success": true,
		},
	}
}
