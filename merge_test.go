package go_fanin_fanout_example

import (
	"testing"
)

func TestMergeChannel(t *testing.T) {
	inChan := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			inChan <- i
		}
		close(inChan)
	}()

	fanOutChannel := func(inputChannel <-chan int) <-chan int {
		fanOutChan := make(chan int)
		go func() {
			for in := range inputChannel {
				fanOutChan <- in
			}
			close(fanOutChan)
		}()
		return fanOutChan
	}

	var fanOuts []<-chan int
	for i := 0; i < 100; i++ {
		fanOuts = append(fanOuts, fanOutChannel(inChan))
	}

	var results []int
	for m := range MergeChannel(fanOuts...) {
		results = append(results, m)
	}

	if len(results) != 100 {
		t.Errorf("results size must be 100")
	}
}
