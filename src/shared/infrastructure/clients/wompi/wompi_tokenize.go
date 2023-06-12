package wompi_client

import (
	"encoding/json"
	"log"

	wompi_products "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi/products"
)

type TokenizeParams struct {
	Type   string
	Url    string
	Params any
}

type TokenizeOption func(*TokenizeParams)

func TokenizeCreditCardOption(creditCard wompi_products.CreditCard) TokenizeOption {
	return func(tokenizeParams *TokenizeParams) {
		tokenizeParams.Url = "v1/tokens/cards"
		tokenizeParams.Params = creditCard
	}
}

func TokenizeNequiOption(nequi wompi_products.Nequi) TokenizeOption {
	return func(tokenizeParams *TokenizeParams) {
		tokenizeParams.Url = "v1/tokens/nequi"
		tokenizeParams.Params = nequi
	}
}

func NewTokenizeParams(tokenizeOptions ...TokenizeOption) (tokenizeParmas TokenizeParams) {
	for _, tokenizeOption := range tokenizeOptions {
		tokenizeOption(&tokenizeParmas)
	}

	return
}

func (w *WompiClient) TokenizeCreditCard(tokenizeParams TokenizeParams) CreditCardTokenizeResponseData {
	_, response := w.Fetch(FetchParams{
		Method:        "POST",
		Url:           tokenizeParams.Url,
		UsePrivateKey: false,
		Body:          tokenizeParams.Params,
	})

	var responseBody CreditCardTokenizeResponseData
	err := json.Unmarshal(response, &responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	return responseBody
}
