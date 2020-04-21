package main

import (
	"os"

	"github.com/cnbattle/aitranslate/utils"
)

// 还原初始化icon图片
func initFile() {
	home, err := utils.Home()
	if err != nil {
		panic(err)
	}
	path = home + "/.aitranslate/"
	_ = os.Mkdir(path, os.ModePerm)
	if !utils.Exists(path + "images/logo.png") {
		_ = RestoreAssets(path, "images/logo.png")
		_ = RestoreAssets(path, "images/warning.png")
	}
}

// getImagePath 获取图片地址
func getImagePath(image string) string {
	switch image {
	case "logo":
		return path + "images/logo.png"
	default:
		return path + "images/warning.png"
	}
}
