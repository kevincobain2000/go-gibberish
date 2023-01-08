package gibberish

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/go-shiori/dom"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/markusmobius/go-trafilatura"
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

func ExtractReadableText(URL string) (string, error) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	parsedURL, err := url.ParseRequestURI(URL)
	if err != nil {
		return "", err
	}

	// Fetch article
	resp, err := httpClient.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Extract content
	opts := trafilatura.Options{
		IncludeImages: false,
		IncludeLinks:  false,
		OriginalURL:   parsedURL,
	}

	result, err := trafilatura.Extract(resp.Body, opts)
	if err != nil {
		return "", err
	}

	readableText := strip.StripTags(dom.OuterHTML(result.ContentNode))

	return readableText, nil
}

func IsURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return false
	}
	return u.IsAbs()
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
