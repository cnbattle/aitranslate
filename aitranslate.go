package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/cnbattle/aitranslate/whatlang"
)

var (
	channel string
	text    string
	path    string
	once    bool
	Version string
	v       bool
)

func init() {
	initFile()
	flag.StringVar(&channel, "c", "Google", "translate channel:  Google or YouDao.")
	flag.BoolVar(&once, "once", false, "run once, default auto translate for clipboard english content.")
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
			time.Sleep(time.Millisecond * 150)
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
