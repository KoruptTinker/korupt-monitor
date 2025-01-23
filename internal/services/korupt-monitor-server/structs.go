package korupt_monitor_server

import httpClient "github.com/KoruptTinker/korupt-monitor/internal/core/http_client"

type Service struct {
	httpClient.BaseExternal
}

type RecordClicksRequest struct {
	LeftClickCount  int `json:"left_click_count"`
	RightClickCount int `json:"right_click_count"`
}

type RecordKeypressRequest struct {
	KeypressCount int `json:"key_press_count"`
}
