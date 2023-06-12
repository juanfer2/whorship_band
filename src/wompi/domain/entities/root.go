package wompi_entities

type TypeTransaction uint

const (
	PSE TypeTransaction = iota
	CREDIT_CARD
	NEQUI
)

func (t TypeTransaction) String() string {
	switch t {
	case PSE:
		return "PSE"
	case CREDIT_CARD:
		return "CARD"
	case NEQUI:
		return "NEQUI"
	}
	return "unknown"
}

type CustomerData struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
}

type Payment struct {
	Type            TypeTransaction
	AmountInCents   int
	Currency        string
	AcceptanceToken string
	CustomerEmail   string
	Reference       string
	CustomerData    CustomerData
	PaymentMethod   any
}

func NewPayment(
	paramsPaymentBody Payment,
	paymentOptions ...PaymentOption,
) Payment {
	paymentBody := Payment{
		AmountInCents:   paramsPaymentBody.AmountInCents,
		Currency:        paramsPaymentBody.Currency,
		AcceptanceToken: paramsPaymentBody.AcceptanceToken,
		CustomerEmail:   paramsPaymentBody.CustomerEmail,
		Reference:       paramsPaymentBody.Reference,
		CustomerData:    paramsPaymentBody.CustomerData,
	}

	for _, paymentOption := range paymentOptions {
		paymentOption(&paymentBody)
	}

	return paymentBody
}

func (p *Payment) TypeTransaction() string {
	return p.Type.String()
}

func (p *Payment) IsNequi() bool {
	return p.Type == NEQUI
}

func (p *Payment) IsCreditCard() bool {
	return p.Type == CREDIT_CARD
}

func (p *Payment) IsPSE() bool {
	return p.Type == PSE
}

type PaymentOption func(*Payment)

func PaymentWithCreditCard(paymentMethodCreditCard CreditCard) PaymentOption {
	return func(p *Payment) {
		p.Type = CREDIT_CARD
		p.PaymentMethod = paymentMethodCreditCard
	}
}

func PaymentWithPSE(paymentMethodPSE PSEProduct) PaymentOption {
	return func(p *Payment) {
		p.Type = PSE
		p.PaymentMethod = paymentMethodPSE
	}
}

func PaymentWithNequi(paymentMethodPSE Nequi) PaymentOption {
	return func(p *Payment) {
		p.Type = NEQUI
		p.PaymentMethod = paymentMethodPSE
	}
}
