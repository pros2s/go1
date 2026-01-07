package pay

import "fmt"

type Payment struct {
	Desc    string `json:"description"`
	Paid    int    `json:"paid"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (p Payment) Println() {
	fmt.Println("Desc: ", p.Desc)
	fmt.Println("Name: ", p.Name)
	fmt.Println("Surname: ", p.Surname)
	fmt.Println("Paid: ", p.Paid)
	fmt.Println("===========")
}

type PaymentResponse struct {
	Payment
	Time     int64 `json:"time"`
	Sl       []int `json:"slice"`
	ActMoney int   `json:"actmoney"`
}
