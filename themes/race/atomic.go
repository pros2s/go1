package race

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num atomic.Int64

func addAtomicInt(wg *sync.WaitGroup) {
	defer wg.Done()

	for range 1000 {
		num.Add(1)
	}
}

func TestAtomic() {
	wg := sync.WaitGroup{}

	for range 5 {
		wg.Add(1)
		go addAtomicInt(&wg)
	}

	wg.Wait()
	fmt.Println(num.Load())
}
