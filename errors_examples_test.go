package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
	"io"
	"math/rand"
)

func ExampleErrIsAnyOf() {
	funcWithRandomError := func() error {
		errs := []error{
			io.EOF,
			io.ErrNoProgress,
			io.ErrClosedPipe,
		}

		return errs[rand.Intn(len(errs))]
	}

	err := funcWithRandomError()
	// Instead of switch/case you can use:
	fmt.Println(just.ErrIsAnyOf(err, io.EOF, io.ErrClosedPipe, io.ErrNoProgress))
	// Output: true
}

func ExampleErrAs() {
	err := fmt.Errorf("problem: %w", customErr{reason: 13})

	e, ok := just.ErrAs[customErr](err)
	fmt.Printf("%#v, %t", e, ok)
	// Output: just_test.customErr{reason:13}, true
}
