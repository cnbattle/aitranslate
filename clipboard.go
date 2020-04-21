package main

import (
	"github.com/atotto/clipboard"
	"strings"
)

// getClipboardString 获取粘贴版内容 并去除换行符号
func getClipboardString() string {
	newText, _ := clipboard.ReadAll()
	newText = strings.Replace(newText, "\n", "", -1)
	return newText
}
