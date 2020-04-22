package main

import (
	"github.com/cnbattle/aitranslate/translate"
	"github.com/cnbattle/aitranslate/whatlang"
	"github.com/gen2brain/beeep"
	"strings"
	"time"
)

// translates  Translate and  Notify
func translates() {

	text = getClipboardString()
	for {
		time.Sleep(time.Millisecond * time.Duration(monitoringInterval))
		newText := getClipboardString()
		if !whatlang.IsEnglish(newText) {
			continue
		}
		if !strings.EqualFold(text, newText) {
			text = newText
			translateText, err := translate.Trans(channel, text)
			if err != nil {
				beeep.Alert("Error", err.Error(), getImagePath("warning"))
				return
			}
			beeep.Alert(channel+" Translate", translateText, getImagePath("logo"))
		}
	}
}
