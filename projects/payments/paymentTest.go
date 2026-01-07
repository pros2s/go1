package payments

import (
	"go1/projects/payments/methods"

	"github.com/k0kubun/pp"
)

func PaymentsTest() {
	method := methods.NewPaypal()
	module := NewPaymentModule(method)

	firstBuyID := module.Pay("Buy 1", 10)
	module.Pay("Buy 2", 20)

	module.CancelByID(firstBuyID)
	pp.Println(module.InfoByID(firstBuyID))

	pp.Println(module.Info())
}
