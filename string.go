package just

import (
	"unicode/utf8"
)

// StrCharCount returns rune count in string.
func StrCharCount(s string) int {
	return utf8.RuneCountInString(s)
}

// StrSplitByChars returns slice of runes in string.
func StrSplitByChars(s string) []rune {
	chars := make([]rune, 0, len(s))
	var idx int
	buf := []byte(s)
	for {
		char, size := utf8.DecodeRune(buf[idx:])
		if size == 0 {
			break
		}

		chars = append(chars, char)
		idx += size
	}

	return chars
}

// StrGetFirst return string contains first N valid chars.
func StrGetFirst(s string, n int) string {
	if n <= 0 {
		return ""
	}

	var idx int
	for n > 0 {
		_, size := utf8.DecodeRune([]byte(s)[idx:])
		if size == 0 {
			break
		}

		idx += size
		n--
	}

	return s[:idx]
}
