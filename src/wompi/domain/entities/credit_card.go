package wompi_entities

type CreditCard struct {
	Number     string `json:"number"`
	ExpMonth   string `json:"exp_month"`
	ExpYear    string `json:"exp_year"`
	CVC        string `json:"cvc"`
	CardHolder string `json:"card_holder"`
	Type       string `json:"name"`
}

type CreditCardSource struct {
	Type         string `json:"type"`
	Installments int    `json:"installments"`
	Token        string `json:"token"`
}

type CreditCardSourcePayment struct {
	Payment
	CreditCardSource
}

func NewCreditCardSourcePayment(
	payment Payment, creditCardSource CreditCardSource,
) *CreditCardSourcePayment {
	return &CreditCardSourcePayment{Payment: payment, CreditCardSource: creditCardSource}
}

type CreditCardPayment struct {
	Payment
	CreditCard
}

func NewCreditCardPayment(payment Payment, creditCard CreditCard) *CreditCardPayment {
	return &CreditCardPayment{Payment: payment, CreditCard: creditCard}
}
