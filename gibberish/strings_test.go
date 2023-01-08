package gibberish

import (
	"testing"

	"gotest.tools/assert"
)

func TestRemoveNumbers(t *testing.T) {
	tests := []struct {
		sentence string
		want     string
	}{
		{
			sentence: "18",
			want:     "  ",
		},
		{
			sentence: "This is a sentence",
			want:     "This is a sentence",
		},
		{
			sentence: "sentence 😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊",
			want:     "sentence 😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊",
		},
		{
			sentence: "asdkljas123 dlkasjd isakljd",
			want:     "asdkljas    dlkasjd isakljd",
		},
		{
			sentence: "sentence123 123 dlkasjd dlkasjd isakljd",
			want:     "sentence        dlkasjd dlkasjd isakljd",
		},
	}
	for _, test := range tests {
		t.Run(test.sentence, func(t *testing.T) {
			actual := RemoveNumbers(test.sentence)
			assert.Equal(t, test.want, actual)
		})
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		sentence string
		want     bool
	}{
		{
			sentence: "http://google.com",
			want:     true,
		},
		{
			sentence: "https://google.com/test?q=123",
			want:     true,
		},
	}
	for _, test := range tests {
		t.Run(test.sentence, func(t *testing.T) {
			actual := IsURL(test.sentence)
			assert.Equal(t, test.want, actual)
		})
	}
}
