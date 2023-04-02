package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
)

func ExampleMaxOr() {
	values := []int{10, 20, 30}
	maxValue := just.MaxOr(999, values...)
	// This will print 30, because you have non-empty slice.
	fmt.Println(maxValue)
	// Output: 30
}

func ExampleMin() {
	values := []int{10, 20, 30}
	minValue := just.Min(values...)
	fmt.Println(minValue)
	// Output: 10
}

func ExampleSum() {
	values := []int{10, 20, 30}
	minValue := just.Sum(values...)
	fmt.Println(minValue)
	// Output: 60
}
