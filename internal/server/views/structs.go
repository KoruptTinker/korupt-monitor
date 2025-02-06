package views

type WeeklyDataResponse struct {
	Success bool         `json:"success"`
	Data    []WeeklyData `json:"data"`
}

type WeeklyData struct {
	Date            string `json:"date"`
	KeyPressCount   int    `json:"keyPresses"`
	LeftClickCount  int    `json:"leftClicks"`
	RightClickCount int    `json:"rightClicks"`
}
