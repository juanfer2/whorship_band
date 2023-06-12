package wompi_client

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

type SourcePaymentParams struct {
	Type            string `json:"type"`
	Token           string `json:"token"`
	CustomerEmail   string `json:"customer_email"`
	AcceptanceToken string `json:"acceptance_token"`
}

type PublicData struct {
	Bin            string    `json:"bin"`
	LastFour       string    `json:"last_four"`
	ExpMonth       string    `json:"exp_month"`
	ExpYear        string    `json:"exp_year"`
	CardHolder     string    `json:"card_holder"`
	ValidityEndsAt time.Time `json:"validity_ends_at"`
	Type           string    `json:"type"`
}

type SourcePaymentData struct {
	ID            int        `json:"id"`
	PublicData    PublicData `json:"public_data"`
	Token         string     `json:"token"`
	Type          string     `json:"type"`
	Status        string     `json:"status"`
	CustomerEmail string     `json:"customer_email"`
}

type SourcePaymentResponse struct {
	SourcePaymentData SourcePaymentData `json:"data"`
	Meta              any               `json:"meta"`
}

func (w *WompiClient) SourceCreditCard(sourcePaymentParams SourcePaymentParams) SourcePaymentResponse {
	sourcePaymentParams.AcceptanceToken = w.acceptanceToken()
	_, response := w.Fetch(FetchParams{
		Method:        "POST",
		Url:           "v1/payment_sources",
		UsePrivateKey: true,
		Body:          sourcePaymentParams,
	})

	var errresponseBody any
	json.Unmarshal(response, &errresponseBody)

	var responseBody SourcePaymentResponse
	err := json.Unmarshal(response, &responseBody)
	if err != nil {
		fmt.Println("Error")
		utilities.PrintJson(responseBody)
		log.Fatalln(err)
	}

	utilities.PrintJson(errresponseBody)
	return responseBody
}
