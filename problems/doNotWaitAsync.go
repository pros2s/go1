package problems

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func getDiscount() float64 {
	time.Sleep(2 * time.Second)
	return 12.0
}

func TestDoNotWaitAsync() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelCtx()

	discount, err := checkTimeout(ctx, getDiscount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your discount: %v", discount)
}

func checkTimeout(ctx context.Context, fun func() float64) (float64, error) {
	ch := make(chan float64)

	go func() {
		ch <- fun()
	}()

	select {
	case <-ctx.Done():
		return 0, errors.New("context error")
	case v, ok := <-ch:
		if !ok {
			return 0, errors.New("error with read channel")
		}

		return v, nil
	}
}
