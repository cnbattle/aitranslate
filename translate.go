package main

import (
	"github.com/cnbattle/aitranslate/translate"
	"github.com/gen2brain/beeep"
	"strings"
	"time"
)

// translates  Translate and  Notify
func translates() {
	emptyTimes := 0
	beeep.Alert("", "自动翻译运行中...", "")
	text = getClipboardString()
	for {
		time.Sleep(time.Millisecond * time.Duration(monitoringInterval))
		newText := getClipboardString()
		if strings.EqualFold("", newText) {
			emptyTimes++
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
		if emptyTimes == 30 {
			beeep.Alert("Error", "多次未检查到内容，请确定依赖包已经安装", getImagePath("warning"))
		}
	}
}
