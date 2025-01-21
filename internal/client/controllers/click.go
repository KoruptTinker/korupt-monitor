package controllers

import (
	hook "github.com/robotn/gohook"
)

func (client *ClientController) RecordClick() {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		switch e.Button {
		case hook.MouseMap["left"]:
			client.UserClicks.LeftClicks++
		case hook.MouseMap["right"]:
			client.UserClicks.RightClicks++
		}
	})
}
