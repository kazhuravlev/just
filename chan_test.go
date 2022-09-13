package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
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
