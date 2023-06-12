package wompi_serializers

import wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"

type CreditCardSerializer struct {
	Number     string `json:"number"`
	ExpMonth   string `json:"expMonth"`
	ExpYear    string `json:"expYear"`
	CVC        string `json:"cvc"`
	CardHolder string `json:"cardHolder"`
	Type       string `json:"type"`
}

func NewCreditCardSerializer(creditCard wompi_entities.CreditCard) *CreditCardSerializer {
	return &CreditCardSerializer{
		Number:     creditCard.Number,
		ExpMonth:   creditCard.ExpMonth,
		ExpYear:    creditCard.ExpYear,
		CVC:        creditCard.CVC,
		CardHolder: creditCard.CardHolder,
		Type:       creditCard.Type,
	}
}
