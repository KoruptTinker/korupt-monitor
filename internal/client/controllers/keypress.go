package controllers

import (
	hook "github.com/robotn/gohook"
)

func (client *ClientController) RecordKeyPress() {
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		client.UserKeyPresses.Lock.Lock()
		client.UserKeyPresses.KeyPresses++
		client.UserKeyPresses.Lock.Unlock()
	})
}
