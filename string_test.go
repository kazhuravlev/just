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

func TestStrGetFirst(t *testing.T) {
	t.Parallel()

	table := []struct {
		name string
		in   string
		inN  int
		exp  string
	}{
		{
			name: "empty_input_empty_output",
			in:   "",
			inN:  0,
			exp:  "",
		},
		{
			name: "empty_input_empty_output2",
			in:   "hello",
			inN:  0,
			exp:  "",
		},
		{
			name: "negative_in",
			in:   "hello",
			inN:  -10,
			exp:  "",
		},
		{
			name: "overflow_in",
			in:   "hello",
			inN:  10,
			exp:  "hello",
		},
		{
			name: "latin_chars",
			in:   "hello",
			inN:  4,
			exp:  "hell",
		},
		{
			name: "cyrillic_chars",
			in:   "тесто",
			inN:  4,
			exp:  "тест",
		},
		{
			name: "emoji",
			in:   "🍕",
			inN:  1,
			exp:  "🍕",
		},
		{
			name: "japanese_chars",
			in:   "空母",
			inN:  2,
			exp:  "空母",
		},
	}
	for i := range table {
		row := table[i]
		t.Run(row.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, row.exp, just.StrGetFirst(row.in, row.inN))
		})
	}
}

func BenchmarkStrGetFirst(b *testing.B) {
	str := "Hey, this is a test string for experiment. It length 64 symbols."

	for i := 0; i < b.N; i++ {
		_ = just.StrGetFirst(str, 32)
	}
}
