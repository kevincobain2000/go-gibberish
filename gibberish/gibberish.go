package gibberish

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
)

type Gibberish struct {
	confidenceThreshhold float64
	IsGibberish          bool
	Confidence           float64
}

//go:embed en_words.txt
var enWordsTxt []byte

//go:embed en_cities.txt
var enCitiesTxt []byte

var EnglishWords []string

const (
	defaultConfidenceThreshhold = 0.75
)

func init() {
	var once sync.Once
	once.Do(LoadEnglish)
}

func NewGibberish() *Gibberish {
	return &Gibberish{
		confidenceThreshhold: defaultConfidenceThreshhold,
	}
}

// SetConfidenceThreshhold sets the confidence threshhold
// for the gibberish detector. The default is 0.5
// 0.0 to 1.0
func (j *Gibberish) SetConfidenceThreshhold(confidenceThreshhold float64) *Gibberish {
	j.confidenceThreshhold = confidenceThreshhold
	return j
}

func (j *Gibberish) Detect(raw string) *Gibberish {
	j.Confidence = j.CalculateConfidence(raw)
	j.IsGibberish = j.Confidence <= j.confidenceThreshhold
	return j
}

func (j *Gibberish) CalculateConfidence(raw string) float64 {
	raw = RemovePunctuation(raw)
	raw = RemoveNumbers(raw)
	raw = RemoveEmojis(raw)
	raw = RemoveEmailAddress(raw)
	raw = RemoveURL(raw)

	// Tokens. String to array
	words := strings.Fields(raw)

	remainingWordsCount := len(words)

	positiveCount := 0

	for _, word := range words {
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)

		if len(word) == 0 ||
			!j.passDictionaryChallenge(word) {
			continue
		}
		positiveCount++

	}

	if remainingWordsCount == 0 {
		return 100.0
	}

	confidence := float64(positiveCount) / float64(remainingWordsCount)
	return confidence
}

func (j *Gibberish) passDictionaryChallenge(word string) bool {
	return contains(EnglishWords, word)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func LoadEnglish() {
	fmt.Println("loading file english dictionary once....")

	lines := strings.Split(string(enWordsTxt), "\n")
	for _, line := range lines {
		word := strings.ToLower(line)
		EnglishWords = append(EnglishWords, word)
	}

	lines = strings.Split(string(enCitiesTxt), "\n")
	for _, line := range lines {
		word := strings.ToLower(line)
		EnglishWords = append(EnglishWords, word)
	}
}
