package views

import (
	"sort"

	"github.com/KoruptTinker/korupt-monitor/internal/models"
)

func RenderDisplayResponse(clickData []models.ClickWeeklyAggregationData, keyPressData []models.KeypressWeeklyAggregationData) WeeklyDataResponse {
	dataObj := []WeeklyData{}
	sort.Slice(clickData, func(i int, j int) bool {
		return clickData[i].ID.After(clickData[j].ID)
	})

	sort.Slice(keyPressData, func(i int, j int) bool {
		return keyPressData[i].ID.After(keyPressData[j].ID)
	})

	for day := range len(clickData) {
		data := WeeklyData{
			Date:            clickData[day].ID.Format("Jan 02, 2006"),
			LeftClickCount:  clickData[day].LeftClickCount,
			RightClickCount: clickData[day].RightClickCount,
			KeyPressCount:   keyPressData[day].Keypresses,
		}

		dataObj = append(dataObj, data)
	}

	return WeeklyDataResponse{
		Success: true,
		Data:    dataObj,
	}
}
