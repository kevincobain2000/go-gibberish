package gibberish

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"gotest.tools/assert"
)

func TestBasic(t *testing.T) {

	tests := []struct {
		sentence string
		want     bool
	}{
		{
			sentence: "1.8",
			want:     false,
		},
		{
			sentence: "This is a sentence",
			want:     false,
		},
		{
			sentence: "sentence 😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊",
			want:     false,
		},
		{
			sentence: "asdkljas dlkasjd isakljd",
			want:     true,
		},
		{
			sentence: "sentence dlkasjd dlkasjd isakljd",
			want:     true,
		},
	}
	for _, test := range tests {
		t.Run(test.sentence, func(t *testing.T) {
			j := NewGibberish().Detect(test.sentence)
			assert.Equal(t, test.want, j.IsGibberish)
		})
	}
}

func TestNotGibberishes(t *testing.T) {
	for i := 0; i < 100; i++ {
		sentence := gofakeit.Sentence(10)
		t.Run(sentence, func(t *testing.T) {
			j := NewGibberish().Detect(sentence)
			assert.Equal(t, false, j.IsGibberish)
		})

		sentence = gofakeit.Phone()
		t.Run(sentence, func(t *testing.T) {
			j := NewGibberish().Detect(sentence)
			assert.Equal(t, false, j.IsGibberish)
		})
	}
}

// improve this test
// func TestNotGibberishesAccuracy(t *testing.T) {
// 	isNotGubberishCount := 0
// 	for i := 0; i < 100; i++ {
// 		sentence := gofakeit.Question()
// 		t.Run(sentence, func(t *testing.T) {
// 			j := NewGibberish().Detect(sentence)
// 			if !j.IsGibberish {
// 				isNotGubberishCount++
// 			}
// 		})
// 	}
// 	accuracy := float64(isNotGubberishCount) / 100
// 	assert.Equal(t, true, accuracy > 0.90)
// }

func TestGibberishesAccuracy(t *testing.T) {
	isGubberishCount := 0
	for i := 0; i < 100; i++ {
		sentence := gofakeit.LoremIpsumSentence(10)
		t.Run(sentence, func(t *testing.T) {
			j := NewGibberish().Detect(sentence)
			if j.IsGibberish {
				isGubberishCount++
			}
		})
	}
	accuracy := float64(isGubberishCount) / 100
	assert.Equal(t, true, accuracy > 0.95)

	isGubberishCount = 0
	for i := 0; i < 100; i++ {
		sentence := gofakeit.LoremIpsumSentence(20)
		t.Run(sentence, func(t *testing.T) {
			j := NewGibberish().Detect(sentence)
			if j.IsGibberish {
				isGubberishCount++
			}
		})
	}
	accuracy = float64(isGubberishCount) / 100
	assert.Equal(t, true, accuracy > 0.95)
}

func BenchmarkMyFunction(b *testing.B) {
	// Run the function b.N times
	for n := 0; n < b.N; n++ {
		sentence := gofakeit.Sentence(10)
		NewGibberish().Detect(sentence)
	}
}
