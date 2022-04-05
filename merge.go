package go_fanin_fanout_example

import (
	"sync"
)

//MergeChannel merges channels into a single channel
func MergeChannel[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	multiplex := func(c <-chan T) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range channels {
		wg.Add(1)
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
