package whatlang

import (
	"strings"

	"github.com/abadojack/whatlanggo"
)

// IsEnglish  IsEnglish
func IsEnglish(text string) bool {
	return strings.EqualFold("English", what(text))
}

// what what lang
func what(text string) string {
	options := whatlanggo.Options{
		Whitelist: map[whatlanggo.Lang]bool{
			whatlanggo.Eng: true,
		},
	}
	info := whatlanggo.DetectWithOptions(text, options)
	return info.Lang.String()
}
