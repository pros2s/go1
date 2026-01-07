package goroutines

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/k0kubun/pp"
)

var iterations = 10

func goroutine(channel chan float64, num float64) {
	pp.Printf("Goroutine with number %v\n", num)

	channel <- num
}

func TextGoroutine() {
	result := 0.0
	channel := make(chan float64)

	pp.Printf("Start\n")
	start := time.Now()

	random := rand.Float64()
	for i := range iterations {
		go goroutine(channel, float64(i)*random)
	}

	for range iterations {
		result += <-channel
	}

	pp.Printf("Result: %v\n", result)
	fmt.Printf("Finish with time %v\n", time.Since(start))
}
