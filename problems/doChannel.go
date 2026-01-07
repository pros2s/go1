package problems

import (
	"fmt"
	"time"
)

func doSmt() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	return ch
}

func TestDoSmt() {
	start := time.Now()

	// 2 seconds
	ch1, ch2 := doSmt(), doSmt() // start to funcs immediately
	a, b := <-ch1, <-ch2

	// 4 seconds
	// a, b := <-doSmt(), <-doSmt() // read from each funcs consistently

	fmt.Println(a, b)
	fmt.Println(time.Since(start).Seconds())
}
