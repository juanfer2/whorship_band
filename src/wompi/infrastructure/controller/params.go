package wompi_controller

import (
	"github.com/gofiber/fiber/v2"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
)

type SourceCreditCardParams struct {
	Number        string `json:"number"`
	ExpMonth      string `json:"expMonth"`
	ExpYear       string `json:"expYear"`
	CVC           string `json:"cvc"`
	CardHolder    string `json:"cardHolder"`
	CustomerEmail string `json:"customerEmail"`
}

type CreditCardParams struct {
	CreditCard    wompi_entities.CreditCard
	CustomerEmail string
}

func getParamsCreditCard(c *fiber.Ctx) (CreditCardParams, error) {
	var sourceCreditCardParams SourceCreditCardParams
	if err := c.BodyParser(&sourceCreditCardParams); err != nil {
		return CreditCardParams{}, err
	}

	return CreditCardParams{
		CreditCard: wompi_entities.CreditCard{
			Number:     sourceCreditCardParams.Number,
			ExpMonth:   sourceCreditCardParams.ExpMonth,
			ExpYear:    sourceCreditCardParams.ExpYear,
			CVC:        sourceCreditCardParams.CVC,
			CardHolder: sourceCreditCardParams.CardHolder,
		},
		CustomerEmail: sourceCreditCardParams.CustomerEmail,
	}, nil
}

/*
type Person struct {
	FirstName string `snake:"first_name" camel:"firstName"`
	LastName string `snake:"last_name" camel:"lastName"`
	CurrentAge int `snake:"current_age" camel:"currentAge"`
}

func print() {
	snakeCaseJSON := jsoniter.Config{TagKey: "snake"}.Froze()
	camelCaseJSON := jsoniter.Config{TagKey: "camel"}.Froze()

	p := &Person{"Pepito", "Perez", 32}

	result, _ := snakeCaseJSON.Marshal(p)
	fmt.Println(string(result))

	result, _ = camelCaseJSON.Marshal(p)
	fmt.Println(string(result))
}
*/
