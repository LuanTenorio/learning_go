package srp

type Client struct {
	Name string
	Age  int
}

func (*Client) CalculateExpense() int {
	return 1
}

func (*Client) SendMailToCharging() {

}

func (*Client) SendMailWelcome() {

}

func (*Client) Save() {

}

func (*Client) Delete() {

}
