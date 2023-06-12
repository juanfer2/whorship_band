package wompi_client

import "time"

type Merchant struct {
	ID                     int                 `json:"id"`
	Name                   string              `json:"name"`
	Email                  string              `json:"email"`
	ContactName            string              `json:"contact_name"`
	PhoneNumber            string              `json:"phone_number"`
	Active                 bool                `json:"active"`
	LogoURL                interface{}         `json:"logo_url"`
	LegalName              string              `json:"legal_name"`
	LegalIDType            string              `json:"legal_id_type"`
	LegalID                string              `json:"legal_id"`
	PublicKey              string              `json:"public_key"`
	AcceptedCurrencies     []string            `json:"accepted_currencies"`
	FraudJavascriptKey     interface{}         `json:"fraud_javascript_key"`
	FraudGroups            []interface{}       `json:"fraud_groups"`
	AcceptedPaymentMethods []string            `json:"accepted_payment_methods"`
	PaymentMethods         []PaymentMethod     `json:"payment_methods"`
	PresignedAcceptance    PresignedAcceptance `json:"presigned_acceptance"`
}

type MerchantResponse struct {
	Merchant Merchant `json:"data"`
	Meta     any      `json:"meta"`
}

type PaymentMethod struct {
	Name              string             `json:"name"`
	PaymentProcessors []PaymentProcessor `json:"payment_processors"`
}

type PaymentProcessor struct {
	Name string `json:"name"`
}

type PresignedAcceptance struct {
	AcceptanceToken string `json:"acceptance_token"`
	Permalink       string `json:"permalink"`
	Type            string `json:"type"`
}

type FetchParams struct {
	Url           string
	Method        string
	Body          any
	UsePrivateKey bool
}

type CustomerData struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
}

type PaymentBody struct {
	AmountInCents   int    `json:"amount_in_cents"`
	Currency        string `json:"currency"`
	AcceptanceToken string `json:"acceptance_token"`
	CustomerEmail   string `json:"customer_email"`
	Reference       string `json:"reference"`

	CustomerData  CustomerData `json:"customer_data"`
	PaymentMethod any          `json:"payment_method"`
}

type PaymentMethodPSE struct {
	Type                     string `json:"type"`
	UserType                 int    `json:"user_type"`
	UserLegalIDType          string `json:"user_legal_id_type"`
	UserLegalID              string `json:"user_legal_id"`
	FinancialInstitutionCode string `json:"financial_institution_code"`
	PaymentDescription       string `json:"payment_description"`
}

type PaymentPSE struct {
	AmountInCents   int    `json:"amount_in_cents"`
	Currency        string `json:"currency"`
	AcceptanceToken string `json:"acceptance_token"`
	CustomerEmail   string `json:"customer_email"`
	Reference       string `json:"reference"`

	CustomerData  CustomerData     `json:"customer_data"`
	PaymentMethod PaymentMethodPSE `json:"payment_method"`
}

type PaymentMethodCreditCard struct {
	Type         string `json:"type"`
	Installments int    `json:"installments"`
	Token        string `json:"token"`
}

type PaymentCreditCard struct {
	AmountInCents   int    `json:"amount_in_cents"`
	Currency        string `json:"currency"`
	AcceptanceToken string `json:"acceptance_token"`
	CustomerEmail   string `json:"customer_email"`
	Reference       string `json:"reference"`

	CustomerData  CustomerData            `json:"customer_data"`
	PaymentMethod PaymentMethodCreditCard `json:"payment_method"`
}

type PaymentOption func(*PaymentBody)

type ParamsPaymentBody struct {
	AmountInCents   int
	Currency        string
	AcceptanceToken string
	CustomerEmail   string
	Reference       string
	CustomerData    CustomerData
}

func NewPaymentBody(
	paramsPaymentBody ParamsPaymentBody,
	paymentOptions ...PaymentOption,
) PaymentBody {
	paymentBody := PaymentBody{
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

func PaymentWithCreditCard(paymentMethodCreditCard PaymentMethodCreditCard) PaymentOption {
	return func(p *PaymentBody) {
		p.PaymentMethod = paymentMethodCreditCard
	}
}

func PaymentWithPSE(paymentMethodPSE PaymentMethodPSE) PaymentOption {
	return func(p *PaymentBody) {
		p.PaymentMethod = paymentMethodPSE
	}
}

type CreditCardTokenizeResponse struct {
	Status             string             `json:"status"`
	CreditCardTokenize CreditCardTokenize `json:"data"`
}

type CreditCardTokenize struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Brand          string    `json:"brand"`
	Name           string    `json:"name"`
	LastFour       string    `json:"last_four"`
	Bin            string    `json:"bin"`
	ExpYear        string    `json:"exp_year"`
	ExpMonth       string    `json:"exp_month"`
	CardHolder     string    `json:"card_holder"`
	CreatedWithCvc bool      `json:"created_with_cvc"`
	ExpiresAt      time.Time `json:"expires_at"`
	ValidityEndsAt time.Time `json:"validity_ends_at"`
}

type TransationQueryParms struct {
	StartDate time.Time
	EndDate   time.Time
	Page      uint
	PageSize  uint
}

type TransationCreditCardResponse struct {
	Data struct {
		ID                string    `json:"id"`
		CreatedAt         time.Time `json:"created_at"`
		FinalizedAt       any       `json:"finalized_at"`
		AmountInCents     int       `json:"amount_in_cents"`
		Reference         string    `json:"reference"`
		CustomerEmail     string    `json:"customer_email"`
		Currency          string    `json:"currency"`
		PaymentMethodType string    `json:"payment_method_type"`
		PaymentMethod     struct {
			Type  string `json:"type"`
			Extra struct {
				Bin        string `json:"bin"`
				Name       string `json:"name"`
				Brand      string `json:"brand"`
				ExpYear    string `json:"exp_year"`
				ExpMonth   string `json:"exp_month"`
				LastFour   string `json:"last_four"`
				CardHolder string `json:"card_holder"`
			} `json:"extra"`
			Installments int `json:"installments"`
		} `json:"payment_method"`
		Status          string `json:"status"`
		StatusMessage   any    `json:"status_message"`
		BillingData     any    `json:"billing_data"`
		ShippingAddress any    `json:"shipping_address"`
		RedirectURL     any    `json:"redirect_url"`
		PaymentSourceID int    `json:"payment_source_id"`
		PaymentLinkID   any    `json:"payment_link_id"`
		CustomerData    struct {
			FullName    string `json:"full_name"`
			PhoneNumber string `json:"phone_number"`
		} `json:"customer_data"`
		BillID any   `json:"bill_id"`
		Taxes  []any `json:"taxes"`
	} `json:"data"`
	Meta struct {
	} `json:"meta"`
}
