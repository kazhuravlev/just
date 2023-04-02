package just_test

import (
	"bytes"
	"fmt"
	"github.com/kazhuravlev/just"
	"io"
)

func ExampleBool() {
	fmt.Println(just.Bool(0), just.Bool(1), just.Bool(-1))
	// Output: false true false
}

func ExampleMust() {
	val := just.Must(io.ReadAll(bytes.NewBufferString("this is body!")))
	fmt.Println(string(val))
	// Output:
	// this is body!
}
