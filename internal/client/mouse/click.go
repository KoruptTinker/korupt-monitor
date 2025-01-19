package mouse

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func RecordCountData() {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			fmt.Println("Left Button Clicked")
		} else if e.Button == hook.MouseMap["right"] {
			fmt.Println("Right Button Clicked")
		}
	})
}
