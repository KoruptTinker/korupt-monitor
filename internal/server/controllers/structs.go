package controllers

type BaseResponse struct {
	ResponseCode int
	ResponseData interface{}
}

type RecordClickDataRequest struct {
	LeftClickCount  int `json:"left_click_count"`
	RightClickCount int `json:"right_click_count"`
}

type RecordKeyPressDataRequest struct {
	KeypressCount int `json:"key_press_count"`
}
