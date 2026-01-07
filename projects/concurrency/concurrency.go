package concurrency

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"go1/projects/concurrency/miner"
	"go1/projects/concurrency/postman"
)

func minerTest() int64 {
	minerContext, minerCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		minerCancel()
	}()

	minerChannel := miner.MinerPool(minerContext, 11)
	var coal atomic.Int64

	wg := sync.WaitGroup{}
	wg.Go(func() {
		for v := range minerChannel {
			coal.Add(int64(v))
		}
	})

	wg.Wait()
	return coal.Load()
}

func postmanTest() int {
	postmanContext, postManCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		postManCancel()
	}()

	mailsChannel := postman.PostmanPool(postmanContext, 5)
	var mails []string
	mtx := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Go(func() {
		for v := range mailsChannel {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	})

	wg.Wait()

	mtx.Lock()
	length := len(mails)
	mtx.Unlock()

	return length
}

func ConcurrencyTest() {
	wg := sync.WaitGroup{}

	var minerRes int64
	wg.Go(func() {
		minerRes = minerTest()
	})

	var postmanRes int
	wg.Go(func() {
		postmanRes = postmanTest()
	})

	wg.Wait()

	fmt.Println("Miners: ----", minerRes)
	fmt.Println("Postmen: ----", postmanRes)
}
