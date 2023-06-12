package wompi_repositories

import (
	"time"

	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	wompi_products "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi/products"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
)

type CreditCardTokenizeResponseData struct {
	CreditCardTokenizeStruct CreditCardTokenizeStruct `json:"data"`
	Status                   string                   `json:"status"`
}

type CreditCardTokenizeStruct struct {
	Bin            string    `json:"bin"`
	Brand          string    `json:"brand"`
	CardHolder     string    `json:"card_holder"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedWithCvc bool      `json:"created_with_cvc"`
	ExpMonth       string    `json:"exp_month"`
	ExpYear        string    `json:"exp_year"`
	ExpiresAt      time.Time `json:"expires_at"`
	ID             string    `json:"id"`
	LastFour       string    `json:"last_four"`
	Name           string    `json:"name"`
	ValidityEndsAt time.Time `json:"validity_ends_at"`
}

func (w *WompiRepository) TokenizeCreditCard(
	creditCard wompi_entities.CreditCard, customerEmail string,
) wompi_entities.Product {
	creditCardInput := wompi_client.TokenizeCreditCardOption(wompi_products.CreditCard{
		Number:     creditCard.Number,
		ExpMonth:   creditCard.ExpMonth,
		ExpYear:    creditCard.ExpYear,
		CVC:        creditCard.CVC,
		CardHolder: creditCard.CardHolder,
	})

	response := w.wompiClient.TokenizeCreditCard(wompi_client.NewTokenizeParams(creditCardInput))
	sourceResponse := w.wompiClient.SourceCreditCard(wompi_client.SourcePaymentParams{
		Type:          "CARD",
		Token:         response.CreditCardTokenizeStruct.ID,
		CustomerEmail: customerEmail,
	})

	return wompi_entities.Product{
		Name:     response.CreditCardTokenizeStruct.Brand,
		Bin:      response.CreditCardTokenizeStruct.Bin,
		Type:     wompi_entities.PRODUCT_CREDIT_CARD,
		Token:    response.CreditCardTokenizeStruct.ID,
		Mask:     response.CreditCardTokenizeStruct.Name,
		SourceID: sourceResponse.SourcePaymentData.ID,
	}
}

/*
	//var creditCardTokenizeResponseData CreditCardTokenizeResponseData
		creditCardTokenize := CreditCardTokenizeStruct{
			Bin:            response.CreditCardTokenizeStruct.Bin,
			Brand:          response.CreditCardTokenizeStruct.Brand,
			CardHolder:     response.CreditCardTokenizeStruct.CardHolder,
			CreatedAt:      response.CreditCardTokenizeStruct.CreatedAt,
			CreatedWithCvc: response.CreditCardTokenizeStruct.CreatedWithCvc,
			ExpMonth:       response.CreditCardTokenizeStruct.ExpMonth,
			ExpYear:        response.CreditCardTokenizeStruct.ExpYear,
			ExpiresAt:      response.CreditCardTokenizeStruct.ExpiresAt,
			ID:             response.CreditCardTokenizeStruct.ID,
			LastFour:       response.CreditCardTokenizeStruct.LastFour,
			Name:           response.CreditCardTokenizeStruct.Name,
			ValidityEndsAt: response.CreditCardTokenizeStruct.ValidityEndsAt,
		}
*/
