package gibberish

import (
	_ "embed"
	"regexp"
	"strings"
	"unicode"
)

func RemoveNumbers(input string) string {
	input = strings.Map(func(r rune) rune {
		if unicode.IsNumber(r) {
			return ' '
		}
		return r
	}, input)
	return input
}

func RemovePunctuation(input string) string {
	input = strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return ' '
		}
		return r
	}, input)
	return input
}

func RemoveEmojis(input string) string {
	// Create a regular expression to match special characters and emojis
	re := regexp.MustCompile("[^\x00-\x7F]+")

	// Replace all special characters and emojis with an empty string
	output := re.ReplaceAllString(input, "")
	return output
}

func RemoveEmailAddress(input string) string {
	// Create a regular expression to match special characters and emojis
	re := regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+`)

	// Replace all special characters and emojis with an empty string
	output := re.ReplaceAllString(input, "")
	return output
}

func RemoveURL(input string) string {
	// Create a regular expression to match special characters and emojis
	re := regexp.MustCompile(`https?://\S+`)

	// Replace all special characters and emojis with an empty string
	output := re.ReplaceAllString(input, "")
	return output
}
