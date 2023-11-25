package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrSplitByChars(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   string
		exp  []rune
	}{
		{
			name: "empty_input_empty_output",
			in:   "",
			exp:  []rune{},
		},
		{
			name: "latin_chars",
			in:   "hello",
			exp:  []rune{'h', 'e', 'l', 'l', 'o'},
		},
		{
			name: "cyrillic_chars",
			in:   "тест",
			exp:  []rune{'т', 'е', 'с', 'т'},
		},
		{
			name: "mixed_chars",
			in:   "QЙ",
			exp:  []rune{'Q', 'Й'},
		},
		{
			name: "japanese_chars",
			in:   "空母",
			exp:  []rune{'空', '母'},
		},
	}
	for i := range table {
		row := table[i]
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, row.exp, just.StrSplitByChars(row.in))
			require.Equal(t, len(row.exp), just.StrCharCount(row.in))
		})
	}
}
