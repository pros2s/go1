package goroutines

import (
	"context"
	"fmt"
	"time"
)

func goRoutine(ctx context.Context, ch chan int, ms int) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case ch <- i:
			i++
		}

		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

func readChannel(ch chan int, mes string) {
	for v := range ch {
		fmt.Println(mes, v)
	}
}

func ContextTest() {
	parentCtx, cancelParent := context.WithCancel(context.Background())
	childCtx, cancelChild := context.WithCancel(parentCtx)

	parentCh := make(chan int)
	childCh := make(chan int)

	go readChannel(parentCh, "Parent")
	go readChannel(childCh, "Child")

	go goRoutine(parentCtx, parentCh, 500)
	go goRoutine(childCtx, childCh, 100)

	time.Sleep(1 * time.Second)
	cancelChild()

	time.Sleep(3 * time.Second)
	cancelParent()

	time.Sleep(1 * time.Second)
	fmt.Println("End")
}
