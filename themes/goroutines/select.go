package goroutines

import (
	"fmt"
	"strconv"
	"time"
)

var selectIterations int = 10

func SelectTest() {
	chInt := make(chan int)
	chStr := make(chan string)

	go func() {
		for i := range selectIterations {
			chInt <- i

			time.Sleep(500 * time.Millisecond)
		}

		close(chInt)
	}()

	go func() {
		for i := range selectIterations + 20 {
			chStr <- "Str" + strconv.Itoa(i)

			time.Sleep(100 * time.Millisecond)
		}

		close(chStr)
	}()

	for chInt != nil || chStr != nil {
		select {
		case num, okInt := <-chInt:
			if !okInt {
				chInt = nil
				continue
			}
			fmt.Println(num)
		case str, okStr := <-chStr:
			if !okStr {
				chStr = nil
				continue
			}
			fmt.Println(str)
		}
	}
}
