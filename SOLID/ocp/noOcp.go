package ocp

const (
	debitCard = iota
	creditCard
	ticket
	pix
)

type Finance struct {
}

func (*Finance) ProcessPayment(method int) {
	switch method {
	case debitCard:
		debitCardPayment()
	case creditCard:
		creditCardPayment()
	case ticket:
		ticketPayment()
	case pix:
		pixPayment()
	}
}

func debitCardPayment() {

}

func creditCardPayment() {

}

func ticketPayment() {

}

func pixPayment() {

}
