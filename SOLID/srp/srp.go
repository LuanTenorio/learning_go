package srp

type ClientSrp struct {
	Name string
	Age  int
}

func (*ClientSrp) CalculateExpense() int {
	return 1
}

type Mail interface {
	SendMailToCharging()
	SendMailWelcome()
}

type ClientMail struct {
}

func (*ClientMail) SendMailToCharging() {

}

func (*ClientMail) SendMailWelcome() {

}

type ClientRepository struct {
}

func (*ClientRepository) Save() {

}

func (*ClientRepository) Delete() {

}
