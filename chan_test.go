package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestChanAdapt(t *testing.T) {
	inCh := make(chan int, 3)
	just.ChanPut(inCh, []int{1, 2, 3})

	outCh := just.ChanAdapt(inCh, strconv.Itoa)
	assert.Equal(t, "1", <-outCh)
	assert.Equal(t, "2", <-outCh)
	assert.Equal(t, "3", <-outCh)

	close(inCh)

	_, ok := <-outCh
	assert.False(t, ok, "out channel should be closed at this moment")
}

func TestChanReadN(t *testing.T) {
	const chLen = 10
	const n = chLen / 2
	const msg = "hi"

	messages := just.SliceFillElem(n, msg)

	in := make(chan string, chLen)
	just.ChanPut(in, messages)

	res := just.ChanReadN(in, n)
	require.Equal(t, messages, res)
}
