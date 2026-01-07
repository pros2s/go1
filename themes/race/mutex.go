package race

import (
	"fmt"
	"sync"
)

var slice []int

func addToSlice(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	for range 1000 {
		mtx.Lock()
		slice = append(slice, 1)
		mtx.Unlock()
	}
}

func TestMutex() {
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	for range 5 {
		wg.Add(1)
		go addToSlice(wg, mtx)
	}

	wg.Wait()

	mtx.Lock()
	fmt.Println(len(slice))
	mtx.Unlock()
}
