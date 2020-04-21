package translate

import (
	"errors"
	"net/url"
)

type translateService interface {
	Trans(text string) (string, error)
}

func Trans(scenes, text string) (string, error) {
	if scenes == "" || text == "" {
		return "", errors.New("缺少参数")
	}
	var trans translateService
	switch scenes {
	case "Google":
		trans = googleTranslate{}
	case "YouDao":
		trans = youDaoTranslate{}
	default:
		trans = googleTranslate{}
	}
	return trans.Trans(url.PathEscape(text))
}
