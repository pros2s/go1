package goroutines

import (
	"fmt"
)

type AxiomChannel struct {
	channel chan int
}

func newAxiomChannel(value int) *AxiomChannel {
	newChannel := make(chan int, 1)
	newChannel <- value

	return &AxiomChannel{
		channel: newChannel,
	}
}

func (a *AxiomChannel) getChannel() *chan int {
	return &a.channel
}

func (a *AxiomChannel) setValue(val int) {
	a.channel <- val
}

func (a *AxiomChannel) closeChannel() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("panic with closed channel: %s", fmt.Sprint(rec))
		}
	}()

	close(a.channel)
	return nil
}

var goroutinesCount int = 3

func TestAxioms() {
	channel := newAxiomChannel(15)

	chanVal, ok := <-*channel.getChannel()
	fmt.Println(chanVal, ok)

	go func() {
		for i := range goroutinesCount {
			channel.setValue(i)
		}

		channel.closeChannel()
	}()

	chann := *channel.getChannel()
	channVal, ok1 := <-chann
	fmt.Println(channVal, ok1)

	for val := range chann {
		fmt.Println(val)
	}
}
