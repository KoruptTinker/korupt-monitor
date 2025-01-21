package controllers

type BaseResponse struct {
	ResponseCode int
	ResponseData interface{}
}

type RecordClickDataRequest struct {
	ClickCount int `json:"click_count"`
}

type RecordKeyPressDataRequest struct {
	KeypressCount int `json:"key_press_count"`
}
