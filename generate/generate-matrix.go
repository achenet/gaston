package generate

import "strings"

// GenerateTransitionmatrix a transition matrix from some text
func GenerateTransitionmatrix(text string) [][]float32 {
	text = preProcess(text)
	if len(text) == 0 {
		return nil
	}
	out := make([][]float32, 26)
	for i := range out {
		out[i] = make([]float32, 26)
	}
	i := 0
	runes := []rune(text)
	for i < (len(runes) - 1) {
		out[numberOf[runes[i]]][numberOf[runes[i+1]]] += 1
		i += 1
	}
	l := len([]rune(text)) - 1
	for _, row := range out {
		for i := range row {
			row[i] = row[i] / float32(l)
		}
	}
	return out
}

var numberOf = map[rune]uint8{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}

func preProcess(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "") // TODO use a regexp
	return text
}

func FindFrequencies(text string) []float32 {
	text = preProcess(text)
	out := make([]float32, 26)
	for _, r := range text {
		out[numberOf[r]] += 1
	}
	// normalize
	l := len([]rune(text))
	for i := range out {
		out[i] = out[i] / float32(l)
	}
	return out
}
