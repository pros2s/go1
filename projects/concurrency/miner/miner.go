package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	minerCh chan<- int,
	curMiner int,
	power int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Miner %d with %d --> Finished\n", curMiner, power)
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Miner %d --> %d\n", curMiner, power)
		}

		select {
		case <-ctx.Done():
			fmt.Printf("Miner %d with %d --> Finished\n", curMiner, power)
			return
		case minerCh <- power:
			fmt.Printf("Miner %d --> %d\n", curMiner, power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	wg := &sync.WaitGroup{}
	minerCh := make(chan int)

	for i := range minerCount {
		wg.Add(1)
		go miner(ctx, wg, minerCh, i, i+10)
	}

	go func() {
		wg.Wait()
		close(minerCh)
	}()

	return minerCh
}
