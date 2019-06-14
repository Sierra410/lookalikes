package lookalikes

import (
	"strings"
	"testing"
)

const testString = "Thｅ quｉcк ᖯroｗn ｆօx јumps oѵer thｅ ǀɑzy doɡ 012Ȝ4Ƽ67ȣ9"
const testStringProper = "The quick brown fox jumps over the lazy dog 0123456789"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns a string that consists of 2 characters - space and ^
// spaces where both strings match and ^ where they don't
func diffStr(str1, str2 string) string {
	var (
		s1 = []rune(str1)
		s2 = []rune(str2)

		minl = min(len(s1), len(s2))
		maxl = max(len(s1), len(s2))

		b = strings.Builder{}
	)

	b.Grow(maxl)

	i := 0
	for ; i < minl; i++ {
		if s1[i] == s2[i] {
			b.WriteRune(' ')
		} else {
			b.WriteRune('^')
		}
	}

	for ; i < maxl; i++ {
		b.WriteRune('^')
	}

	return b.String()
}

func TestDecode(t *testing.T) {
	got := Decode(testString)
	if got != testStringProper {
		diff := diffStr(got, testStringProper)
		t.Errorf("\nUnexpected result:\n%s\n%s\n%s", got, diff, testStringProper)
	}
}

func BenchmarkDecodeJunk(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Decode(testString)
	}
}

func BenchmarkDecodeProper(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Decode(testStringProper)
	}
}
