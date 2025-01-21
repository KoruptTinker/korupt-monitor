package controllers

import "sync"

type ClickData struct {
	LeftClicks  int
	RightClicks int
	Lock        *sync.Mutex
}

type KeyboardData struct {
	KeyPresses int
	Lock       *sync.Mutex
}
