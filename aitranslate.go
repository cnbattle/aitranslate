package main

import (
	"flag"
)

var (
	// 翻译渠道
	channel string
	// 待翻译的文本
	text string
	// 通知logo
	path string
	// 监听时间间隔
	monitoringInterval int = 100
)

func init() {
	initFile()
	flag.StringVar(&channel, "c", "Google", "translate channel:  Google or YouDao.")
	flag.Parse()
}

func main() {
	translates()
}
