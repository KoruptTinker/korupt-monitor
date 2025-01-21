package controllers

import (
	hook "github.com/robotn/gohook"
)

func (client *ClientController) RecordClick() {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		switch e.Button {
		case hook.MouseMap["left"]:
			client.UserClicks.Lock.Lock()
			client.UserClicks.LeftClicks++
			client.UserClicks.Lock.Unlock()
		case hook.MouseMap["right"]:
			client.UserClicks.Lock.Lock()
			client.UserClicks.RightClicks++
			client.UserClicks.Lock.Unlock()
		}
	})
}
