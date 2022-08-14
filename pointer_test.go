package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(t *testing.T) {
	a := just.Pointer(10)
	assert.Equal(t, 10, *a)
}
