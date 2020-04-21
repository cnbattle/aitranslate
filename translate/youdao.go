package translate

import (
	"encoding/json"
	"errors"
	"github.com/cnbattle/aitranslate/utils"
)

type youDaoTranslate struct{}

type result struct {
	Type            string              `json:"type"`
	ErrorCode       int                 `json:"errorCode"`
	ElapsedTime     int                 `json:"elapsedTime"`
	TranslateResult [][]translateResult `json:"translateResult"`
}

type translateResult struct {
	Src string `json:"src"`
	Tgt string `json:"tgt"`
}

func (g youDaoTranslate) Trans(text string) (resText string, err error) {
	urlStr := "http://fanyi.youdao.com/translate?&doctype=json&type=EN2ZH_CN&i=" + text
	body, err := utils.HttpGet(urlStr)
	if err != nil {
		return "", err
	}

	var resJson result
	err = json.Unmarshal(body, &resJson)
	if err != nil {
		return "", err
	}

	for _, value := range resJson.TranslateResult[0] {
		resText += value.Tgt
	}
	if len(resText) == 0 {
		return "", errors.New("request api error or empty")
	}
	return resText, nil
}
