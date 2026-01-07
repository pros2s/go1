package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	wg *sync.WaitGroup,
	ctx context.Context,
	ch chan<- string,
	curPostman int,
	letter string,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Postman %d with Letter: %s has finished\n", curPostman, letter)
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Postman %d --> %s\n", curPostman, letter)
		}

		select {
		case <-ctx.Done():
			fmt.Printf("Postman %d with Letter: %s has finished\n", curPostman, letter)
			return
		case ch <- letter:
			fmt.Printf("Postman %d --> %s\n", curPostman, letter)
		}
	}
}

func postManMail(key int) string {
	mailMap := make(map[int]string, 3)

	mailMap[0] = "First letter"
	mailMap[1] = "Second letter"
	mailMap[2] = "Third letter"

	val, ok := mailMap[key]
	if !ok {
		return "Not ok"
	}

	return val
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	wg := &sync.WaitGroup{}
	postmanChannel := make(chan string)

	for i := range postmanCount {
		wg.Add(1)
		go postman(wg, ctx, postmanChannel, i, postManMail(i))
	}

	go func() {
		wg.Wait()
		close(postmanChannel)
	}()

	return postmanChannel
}
