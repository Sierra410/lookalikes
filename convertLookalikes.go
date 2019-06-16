//Package lookalikes consists of a function that converts 𝐬𝐭𝐮𝐟𝐟 𝒍𝒊𝒌𝒆 тЂіѕ
//(popular unicode confusables/lookalikes and other symbols that might be used as letters)
//into proper english text.
//
//Well, most of the times.
package lookalikes

import "strings"

func decodeRune(r rune) rune {
	//No point in testing ASCII
	if r <= 127 {
		return r
	}

	c, ok := lookup[r]

	if ok {
		return c
	}

	return r
}

// Decode "decodes" 𝖘𝖙𝖚𝖋𝖋.
//
// Some symbols that only exist in one form and don't have upper- or lowercase
// variants, and symbols that can represent more than 1 symbol are always
// converted to the first match.
//
// Example: '𝐥' can be 'l', 'L' and '1'. By default it's converted to 'l'.
// This behavior can be changed boringly, by rearranging stuff in the sources.
func Decode(s string) string {
	builder := strings.Builder{}
	builder.Grow(len(s))

	for _, r := range s {
		builder.WriteRune(decodeRune(r))
	}

	return builder.String()
}
