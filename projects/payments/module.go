package payments

import "maps"

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentInfo   map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentInfo:   make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p *PaymentModule) Pay(desc string, usd int) int {
	id := p.paymentMethod.Pay(usd)

	newInfo := PaymentInfo{
		description: desc,
		usd:         usd,
		isCancelled: false,
	}
	p.paymentInfo[id] = newInfo

	return id
}

func (p *PaymentModule) CancelByID(id int) {
	info, ok := p.paymentInfo[id]

	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.isCancelled = true
	p.paymentInfo[id] = info
}

func (p PaymentModule) InfoByID(id int) PaymentInfo {
	info, ok := p.paymentInfo[id]

	if !ok {
		return PaymentInfo{}
	}

	return info
}

func (p PaymentModule) Info() map[int]PaymentInfo {
	copyInfo := make(map[int]PaymentInfo, len(p.paymentInfo))

	maps.Copy(copyInfo, p.paymentInfo)

	return copyInfo
}
