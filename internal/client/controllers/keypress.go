package controllers

import (
	"context"

	hook "github.com/robotn/gohook"
)

func (client *ClientController) RecordKeyPress() {
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		client.UserKeyPresses.Lock.Lock()
		client.UserKeyPresses.KeyPresses++
		client.UserKeyPresses.Lock.Unlock()
	})
}

func (client *ClientController) TransmitKeypressData() error {
	client.UserKeyPresses.Lock.Lock()

	if err := client.External.KoruptMonitorServer.RecordKeypressData(
		context.Background(), client.UserKeyPresses.KeyPresses,
	); err != nil {
		return err
	}

	client.UserKeyPresses.KeyPresses = 0

	client.UserKeyPresses.Lock.Unlock()

	return nil
}
