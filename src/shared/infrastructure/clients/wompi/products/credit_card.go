package wompi_products

type CreditCard struct {
	Number     string `json:"number"`
	ExpMonth   string `json:"exp_month"`
	ExpYear    string `json:"exp_year"`
	CVC        string `json:"cvc"`
	CardHolder string `json:"card_holder"`
}

func (c *CreditCard) CardHolderIsValid() bool {
	return c.CardHolder != ""
}

func (c *CreditCard) NumberIsValid() bool {
	return c.Number != ""
}

func (c *CreditCard) ExpMonthIsValid() bool {
	return c.Number != ""
}

func (c *CreditCard) ExpYearIsValid() bool {
	return c.Number != ""
}

func (c *CreditCard) CVCIsValid() bool {
	return c.Number != ""
}
