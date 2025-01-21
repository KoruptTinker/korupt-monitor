package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) RecordUserClick(ctx *gin.Context) BaseResponse {
	requestData := RecordClickDataRequest{}

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		return BaseResponse{
			ResponseCode: http.StatusBadRequest,
			ResponseData: map[string]string{
				"message": "Failed to parse request body",
				"err":     err.Error(),
			},
		}
	}

	if err := c.DB.InsertClickData(ctx, requestData.ClickCount); err != nil {
		return BaseResponse{
			ResponseCode: http.StatusInternalServerError,
			ResponseData: map[string]string{
				"message": "Something went wrong",
				"err":     err.Error(),
			},
		}
	}

	return BaseResponse{
		ResponseCode: http.StatusOK,
		ResponseData: map[string]string{
			"message": "success",
			"err":     "",
		},
	}
}
