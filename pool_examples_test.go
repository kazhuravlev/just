package just_test

import (
	"bytes"
	"fmt"
	"github.com/kazhuravlev/just"
)

func ExampleNewPool() {
	p := just.NewPool(
		func() *bytes.Buffer { return bytes.NewBuffer(nil) },
		func(buf *bytes.Buffer) { fmt.Println(buf.String()); buf.Reset() },
	)

	buf := p.Get()
	buf.WriteString("Some data")
	// This example showing that Pool will call reset callback
	// on each Pool.Put call.
	p.Put(buf)
	// Output: Some data
}
