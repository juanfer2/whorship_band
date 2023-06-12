package wompi_entities

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

func NewMerchant(
	id int,
	name string,
	email string,
	contactName string,
	phoneNumber string,
	active bool,
	logoURL interface{},
	legalName string,
	legalIDType string,
	legalID string,
	publicKey string,
	acceptedCurrencies []string,
	fraudJavascriptKey interface{},
	fraudGroups []interface{},
	acceptedPaymentMethods []string,
	paymentMethods []PaymentMethod,
	presignedAcceptance PresignedAcceptance,
) *Merchant {
	return &Merchant{
		ID:                     id,
		Name:                   name,
		Email:                  email,
		ContactName:            contactName,
		PhoneNumber:            phoneNumber,
		Active:                 active,
		LogoURL:                logoURL,
		LegalName:              legalName,
		LegalIDType:            legalIDType,
		LegalID:                legalID,
		PublicKey:              publicKey,
		AcceptedCurrencies:     acceptedCurrencies,
		FraudJavascriptKey:     fraudJavascriptKey,
		FraudGroups:            fraudGroups,
		AcceptedPaymentMethods: acceptedPaymentMethods,
		PaymentMethods:         paymentMethods,
		PresignedAcceptance:    presignedAcceptance,
	}
}
