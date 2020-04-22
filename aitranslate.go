package main

import (
	"flag"
	"fmt"
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
	// 是否显示版本号
	v bool
	// 监听时间间隔
	monitoringInterval int
)

func init() {
	initFile()
	flag.StringVar(&channel, "c", "Google", "translate channel:  Google or YouDao.")
	flag.IntVar(&monitoringInterval, "t", 100, "monitoring interval, unit: ms.")
	flag.BoolVar(&v, "v", false, "show version and exit.")
	flag.Parse()
}

func main() {
	if v {
		fmt.Println(Version)
		return
	}
	translates()
}
