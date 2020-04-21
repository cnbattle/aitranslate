package main

import (
	"github.com/gen2brain/beeep"

	"github.com/cnbattle/aitranslate/translate"
)

// translates  Translate and  Notify
func translates() {
	translateText, err := translate.Trans(channel, text)
	if err != nil {
		beeep.Alert("Error", err.Error(), getImagePath("warning"))
		return
	}
	beeep.Notify(channel+" Translate", translateText, getImagePath("logo"))
}
