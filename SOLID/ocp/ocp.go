package ocp

type FinanceOcp struct {
}

type Payment interface {
	Pay()
}

func (*FinanceOcp) ProcessPayment(payment Payment) {
	payment.Pay()
}

type DebitCardPayment struct {
}

func (*DebitCardPayment) Pay() {

}

type CreditCardPayment struct {
}

func (*CreditCardPayment) Pay() {

}

type TicketPayment struct {
}

func (*TicketPayment) Pay() {

}

type PixPayment struct {
}

func (*PixPayment) Pay() {

}
