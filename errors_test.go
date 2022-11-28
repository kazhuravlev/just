package just_test

import (
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
)

func TestErrIsAnyOf(t *testing.T) {
	t.Parallel()

	wrapped := func(e error) error { return fmt.Errorf("wrapped %w", e) }

	table := []struct {
		err      error
		errSlice []error
		exp      bool
	}{
		{
			err:      nil,
			errSlice: nil,
			exp:      false,
		},
		{
			err:      io.EOF,
			errSlice: []error{io.EOF, io.ErrClosedPipe},
			exp:      true,
		},
		{
			err:      wrapped(io.EOF),
			errSlice: []error{io.EOF, io.ErrClosedPipe},
			exp:      true,
		},
		{
			err:      wrapped(io.EOF),
			errSlice: []error{io.ErrClosedPipe},
			exp:      false,
		},
		{
			err:      wrapped(io.EOF),
			errSlice: nil,
			exp:      false,
		},
	}

	for _, row := range table {
		t.Run("", func(t *testing.T) {
			res := just.ErrIsAnyOf(row.err, row.errSlice...)
			assert.Equal(t, row.exp, res)
		})
	}
}

func TestErrIsNotAnyOf(t *testing.T) {
	t.Parallel()

	wrapped := func(e error) error { return fmt.Errorf("wrapped %w", e) }

	table := []struct {
		err      error
		errSlice []error
		exp      bool
	}{
		{
			err:      nil,
			errSlice: nil,
			exp:      true,
		},
		{
			err:      io.EOF,
			errSlice: []error{io.EOF, io.ErrClosedPipe},
			exp:      false,
		},
		{
			err:      wrapped(io.EOF),
			errSlice: []error{io.EOF, io.ErrClosedPipe},
			exp:      false,
		},
		{
			err:      wrapped(io.EOF),
			errSlice: []error{io.ErrClosedPipe},
			exp:      true,
		},
		{
			err:      wrapped(io.EOF),
			errSlice: nil,
			exp:      true,
		},
	}

	for _, row := range table {
		t.Run("", func(t *testing.T) {
			res := just.ErrIsNotAnyOf(row.err, row.errSlice...)
			assert.Equal(t, row.exp, res)
		})
	}
}

type customErr struct {
	reason int
}

func (c customErr) Error() string {
	return strconv.Itoa(c.reason)
}

func TestAsTestAs(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("problem: %w", customErr{reason: 13})

	e, ok := just.As[customErr](err)
	assert.True(t, ok)
	assert.Equal(t, customErr{reason: 13}, e)
}
