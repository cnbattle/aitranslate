package translate

import (
	"errors"
	"github.com/cnbattle/aitranslate/utils"
	"github.com/tidwall/gjson"
	"strconv"
)

type googleTranslate struct{}

func (g googleTranslate) Trans(text string) (str string, err error) {
	urlStr := "http://translate.google.cn/translate_a/single?client=gtx&sl=en&tl=zh-CN&dt=t&q=" + text
	body, err := utils.HttpGet(urlStr)
	if err != nil {
		return "", err
	}

	for i := 0; ; i++ {
		path := "0." + strconv.Itoa(i) + ".0"
		value := gjson.Get(string(body), path)
		if value.Exists() {
			str += value.String()
			continue
		}
		break
	}
	if len(str) == 0 {
		return "", errors.New("request api error")
	}
	return
}
