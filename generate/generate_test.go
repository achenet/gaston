package generate

import "testing"

func TestFindFrequencies(t *testing.T) {
	for text, expect := range map[string][]float32{} {
		got := findFrequencies(text)
		if !areEqual(got, expect) {
			t.Errorf("input: %s, got: %v, expect %v", text, got, expect)
		}
	}
}

func areEqual(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}
