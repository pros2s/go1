package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NewPaypal() Paypal {
	return Paypal{}
}

func (c Paypal) Pay(usd int) int {
	fmt.Println("Paypal pay")

	return rand.Int()
}

func (c Paypal) Cancel(id int) {
	fmt.Printf("Paypal cancelled by id: %d\n", id)
}
