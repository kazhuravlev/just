package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
)

func ExamplePointerUnwrapDefault() {
	someFuncThatReturnsNil := func() *int { return nil }

	value := just.PointerUnwrapDefault(someFuncThatReturnsNil(), 42)
	fmt.Println(value)
	// Output: 42
}
