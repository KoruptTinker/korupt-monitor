package controllers

import (
	"github.com/kbinani/screenshot"
)

func (client *ClientController) CaptureScreen() ([]byte, error) {
	numDisplays := screenshot.NumActiveDisplays()

	if numDisplays > 1 {
		img, err := screenshot.CaptureDisplay(1)
		if err != nil {
			return nil, err
		}
		return img.Pix, nil
	} else {
		img, err := screenshot.CaptureDisplay(0)
		if err != nil {
			return nil, err
		}
		return img.Pix, nil
	}
}
