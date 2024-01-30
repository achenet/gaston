package generate

import "strings"

// Generate a transition matrix from some text
func Generate(text string) [][]float32 {
	text = preProcess(text)
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
	// TODO normalize
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
	return out
}
