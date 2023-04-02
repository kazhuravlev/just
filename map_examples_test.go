package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
)

func ExampleMapGetDefault() {
	m := map[int]int{
		1: 10,
		2: 20,
	}
	val := just.MapGetDefault(m, 3, 42)
	fmt.Println(val)
	// Output: 42
}

func ExampleMapDropKeys() {
	m := map[int]int{
		1: 10,
		2: 20,
	}
	just.MapDropKeys(m, 1, 3)

	fmt.Printf("%#v", m)
	// Output: map[int]int{2:20}
}

func ExampleMapContainsKeysAll() {
	m := map[int]int{
		1: 10,
		2: 20,
	}
	containsAllKeys := just.MapContainsKeysAll(m, []int{1, 2})

	fmt.Println(containsAllKeys)
	// Output: true
}
