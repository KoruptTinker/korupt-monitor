package keyboard

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func RecordCountData() {
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		fmt.Println("Key Pressed")
	})
}
