package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/cnbattle/aitranslate/whatlang"
)

var (
	// 版本号
	Version string
	// 翻译渠道
	channel string
	// 待翻译的文本
	text string
	// 通知logo
	path string
	// 是否单词运行
	once bool
	// 是否显示版本号
	v bool
	// 监听时间间隔
	monitoringInterval int
)

func init() {
	initFile()
	flag.StringVar(&channel, "c", "Google", "translate channel:  Google or YouDao.")
	flag.IntVar(&monitoringInterval, "t", 100, "monitoring interval, unit: ms.")
	flag.BoolVar(&once, "once", false, "run once. (default \"auto translate\")")
	flag.BoolVar(&v, "v", false, "show version and exit.")
	flag.Parse()
}

func main() {
	if v {
		fmt.Println(Version)
		return
	}
	text = getClipboardString()
	if !once {
		for {
			time.Sleep(time.Millisecond * time.Duration(monitoringInterval))
			newText := getClipboardString()
			if !whatlang.IsEnglish(newText) {
				continue
			}
			if !strings.EqualFold(text, newText) {
				text = newText
				translates()
			}
		}
	}
	translates()
}
