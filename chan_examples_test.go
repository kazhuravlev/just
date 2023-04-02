package just_test

import (
	"fmt"
	"github.com/kazhuravlev/just"
	"strconv"
)

func ExampleChanPut() {
	ch := make(chan int, 4)
	just.ChanPut(ch, []int{1, 2, 3, 4})

	resultSlice := just.ChanReadN(ch, 4)
	fmt.Println(resultSlice)
	// Output: [1 2 3 4]
}

func ExampleChanAdapt() {
	intCh := make(chan int, 4)
	just.ChanPut(intCh, []int{1, 2, 3, 4})

	strCh := just.ChanAdapt(intCh, strconv.Itoa)
	fmt.Println(just.ChanReadN(strCh, 4))
	// Output: [1 2 3 4]
}
