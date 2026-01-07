package goroutines

import (
	"fmt"
	"sync"
)

func goWg(wg *sync.WaitGroup, mes string, rang int) {
	defer wg.Done()

	for i := range rang {
		fmt.Println(mes, " go ", i)
	}
}

func WaitGroupTest() {
	wg := &sync.WaitGroup{}

	for i := range 10 {
		wg.Add(1)
		go goWg(wg, fmt.Sprintf("Wg %v", i), i)
	}

	wg.Wait()

	fmt.Println("Finish")
}
