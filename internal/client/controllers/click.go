package controllers

import (
	"context"
	"fmt"

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

func (client *ClientController) TransmitClickData() error {
	client.UserClicks.Lock.Lock()
	if err := client.External.KoruptMonitorServer.RecordClickData(
		context.Background(), client.UserClicks.LeftClicks, client.UserClicks.RightClicks,
	); err != nil {
		fmt.Printf("Encountered error: %v", err.Error())
		return err
	}

	client.UserClicks.LeftClicks = 0
	client.UserClicks.RightClicks = 0

	client.UserClicks.Lock.Unlock()
	return nil
}
