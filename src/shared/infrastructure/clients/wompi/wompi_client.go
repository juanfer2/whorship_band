package wompi_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

type AnyToken interface {
	any | CreditCardTokenizeResponseData
}

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

type WompiClient struct {
	config *ConfigClient
}

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

func NewWompiClient(config *ConfigClient) *WompiClient {
	return &WompiClient{config: config}
}

func (w *WompiClient) GetMerchant() Merchant {
	_, response := w.Fetch(FetchParams{
		Method:        "GET",
		Url:           "v1/merchants/" + w.config.PublicKey,
		UsePrivateKey: true,
	})

	var merchanrResponse MerchantResponse
	json.Unmarshal(response, &merchanrResponse)

	return merchanrResponse.Merchant
}

func (w *WompiClient) CreateTransaction(body PaymentBody) any {
	body.AcceptanceToken = w.acceptanceToken()
	_, response := w.Fetch(FetchParams{
		Method:        "POST",
		Url:           "v1/transactions",
		UsePrivateKey: false,
		Body:          body,
	})

	var responseBody any
	json.Unmarshal(response, &responseBody)

	return responseBody
}

func (w *WompiClient) CreateTransactionWithCreditCard(
	body PaymentBody,
) TransationCreditCardResponse {
	body.AcceptanceToken = w.acceptanceToken()
	_, response := w.Fetch(FetchParams{
		Method:        "POST",
		Url:           "v1/transactions",
		UsePrivateKey: false,
		Body:          body,
	})

	var responseBody TransationCreditCardResponse
	json.Unmarshal(response, &responseBody)

	return responseBody
}

func (w *WompiClient) GetTransactions(queryParams TransationQueryParms) any {
	url := fmt.Sprintf("transactions?from_date='%s'&until_date='%s'&page=%d&page_size=%d",
		queryParams.StartDate, queryParams.EndDate, queryParams.Page, queryParams.PageSize)

	_, response := w.Fetch(FetchParams{
		Method:        "GET",
		Url:           "v1/" + url,
		UsePrivateKey: true,
	})

	var responseBody any
	json.Unmarshal(response, &responseBody)

	return responseBody
}

func (w *WompiClient) Tokenize(tokenizeParams TokenizeParams) any {
	_, response := w.Fetch(FetchParams{
		Method:        "POST",
		Url:           "v1/" + tokenizeParams.Url,
		UsePrivateKey: false,
		Body:          tokenizeParams.Params,
	})

	var responseBody any
	err := json.Unmarshal(response, &responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	return responseBody
}

func (w *WompiClient) acceptanceToken() string {
	return w.GetMerchant().PresignedAcceptance.AcceptanceToken
}

func (w *WompiClient) Fetch(params FetchParams) (error, []byte) {
	var token string
	client := http.Client{}
	var bodyRequest []byte = nil

	if params.Body != nil {
		jsonReq, err := json.Marshal(params.Body)
		if err != nil {
			log.Fatalln(err)
			return err, nil
		}

		bodyRequest = jsonReq
	}

	utilities.PrintJson(params.Body)
	fmt.Println(w.config.Url + "/" + params.Url)

	request, err := http.NewRequest(params.Method, w.config.Url+"/"+params.Url,
		bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
		//os.Exit(1)
	}

	if params.UsePrivateKey {
		token = w.config.PrivateKey
	} else {
		token = w.config.PublicKey
	}

	request.Header = http.Header{
		//"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + token},
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
		//os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	response.Body.Close()

	return err, responseData
}
