package race

import (
	"fmt"
	"sync"
	"time"
)

var slice1 []int

func addToSlice1(wg *sync.WaitGroup, rmtx *sync.RWMutex) {
	defer wg.Done()

	for range 100_000 {
		rmtx.Lock()
		slice1 = append(slice1, 1)
		rmtx.Unlock()
	}
}

func readFromSlice1(wg *sync.WaitGroup, rmtx *sync.RWMutex) {
	defer wg.Done()

	for range 100_000 {
		rmtx.RLock()
		_ = slice1[0]
		rmtx.RUnlock()
	}
}

func TestRwMutex() {
	wg := &sync.WaitGroup{}
	rmtx := &sync.RWMutex{}

	initTime := time.Now()

	for range 100 {
		wg.Add(1)
		go addToSlice1(wg, rmtx)
	}
	wg.Wait()
	fmt.Println("Time: ", time.Since(initTime))

	rTime := time.Now()
	for range 100 {
		wg.Add(1)
		go readFromSlice1(wg, rmtx)
	}

	wg.Wait()
	fmt.Println("R time: ", time.Since(rTime))

	fmt.Println("Length: ", len(slice1))
}
