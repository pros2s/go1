package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Crypto pay")

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Printf("Crypto cancelled by id: %d\n", id)
}
