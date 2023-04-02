package just

// ChanAdapt returns a channel, which will contain adapted messages from
// the source channel. The resulting channel will be closed after the source
// channel is closed.
func ChanAdapt[T, D any](in <-chan T, fn func(T) D) <-chan D {
	ch := make(chan D)
	go func() {
		defer close(ch)

		for elem := range in {
			ch <- fn(elem)
		}
	}()

	return ch
}

// ChanPut will put all elements into the channel synchronously.
func ChanPut[T any](ch chan T, elems []T) {
	for i := range elems {
		ch <- elems[i]
	}
}

// ChanReadN will read N messages from the channel and return the resulting
// slice.
func ChanReadN[T any](ch <-chan T, n int) []T {
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = <-ch
	}

	return res
}
