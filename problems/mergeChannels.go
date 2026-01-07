package problems

import (
	"fmt"
	"sync"
	"time"
)

func gen(data []int, delay time.Duration) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, v := range data {
			time.Sleep(delay)
			out <- v
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	wg := sync.WaitGroup{}
	merged := make(chan int)

	for _, c := range cs {
		wg.Go(func() {
			for v := range c {
				merged <- v
			}
		})
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func TestMergeChannels() {
	c1 := gen([]int{1, 2, 3}, 100*time.Millisecond)
	c2 := gen([]int{4, 5, 6}, 150*time.Millisecond)
	c3 := gen([]int{7, 8, 9}, 50*time.Millisecond)

	merged := merge(c1, c2, c3)

	for v := range merged {
		fmt.Println(v)
	}
}
