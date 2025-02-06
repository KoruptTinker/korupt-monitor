package controllers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/KoruptTinker/korupt-monitor/internal/models"
	"github.com/KoruptTinker/korupt-monitor/internal/server/views"
	"github.com/gin-gonic/gin"
)

func (c *Controller) DisplayWeeklyData(ctx *gin.Context) BaseResponse {
	wg := sync.WaitGroup{}

	clickDataChan := make(chan []models.ClickWeeklyAggregationData, 1)
	keyPressDataChan := make(chan []models.KeypressWeeklyAggregationData, 1)

	wg.Add(2)

	go func() {
		defer wg.Done()
		clickData, err := c.DB.GetWeeklyClickData(ctx.Request.Context())
		if err != nil {
			fmt.Println("Error in click:" + err.Error())
			clickData = []models.ClickWeeklyAggregationData{}
		}
		clickDataChan <- clickData
	}()

	go func() {
		defer wg.Done()

		keyPressData, err := c.DB.GetWeeklyKeypressData(ctx.Request.Context())
		if err != nil {
			fmt.Println("Error encountered:" + err.Error())
			keyPressData = []models.KeypressWeeklyAggregationData{}
		}

		keyPressDataChan <- keyPressData
	}()

	wg.Wait()

	clickData := <-clickDataChan
	keyPressData := <-keyPressDataChan

	response := views.RenderDisplayResponse(clickData, keyPressData)

	return BaseResponse{
		ResponseCode: http.StatusOK,
		ResponseData: response,
	}
}
