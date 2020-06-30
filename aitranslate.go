package main

import (
	"flag"
	"github.com/cnbattle/aitranslate/pkg/icon"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

var (
	// 翻译渠道
	channel string
)

func init() {
	initFile()
	flag.StringVar(&channel, "c", "Google", "translate channel:  Google or YouDao.")
	flag.Parse()
}

func main() {
	go translates()
	systray.Run(onReady, func() {})
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("AiTranslate")

	go func() {
		systray.SetTitle("AiTranslate")
		systray.AddMenuItem("翻译渠道", "").Disable()
		google := systray.AddMenuItem("Google Translate", "Google Translate")
		youdao := systray.AddMenuItem("YouDao Translate", "YouDao Translate")

		systray.AddMenuItem("=====", "").Disable()
		githubUrl := systray.AddMenuItem("Visit Github", "Visit Github")
		quit := systray.AddMenuItem("退出", "Quit the whole app")

		systray.AddSeparator()

		for {
			select {
			case <-google.ClickedCh:
				google.Disable()
				youdao.Enable()
				channel = "Google"
				beeep.Alert("Change", "Use Google Translate now.", getImagePath("logo"))
			case <-youdao.ClickedCh:
				youdao.Disable()
				google.Enable()
				channel = "YouDao"
				beeep.Alert("Change", "Use YouDao Translate now.", getImagePath("logo"))
			case <-githubUrl.ClickedCh:
				open.Run("https://github.com/cnbattle/aitranslate")
			case <-quit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}
