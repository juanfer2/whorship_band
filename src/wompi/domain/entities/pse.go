package wompi_entities

type PSEProduct struct {
	Type                     string `json:"type"`
	UserType                 int    `json:"user_type"`
	UserLegalIDType          string `json:"user_legal_id_type"`
	UserLegalID              string `json:"user_legal_id"`
	FinancialInstitutionCode string `json:"financial_institution_code"`
	PaymentDescription       string `json:"payment_description"`
}

func NewPSEProduct(
	userType int,
	userLegalIDType string,
	userLegalID string,
	financialInstitutionCode string,
	paymentDescription string,
) *PSEProduct {
	return &PSEProduct{
		UserType:                 userType,
		UserLegalIDType:          userLegalIDType,
		UserLegalID:              userLegalID,
		FinancialInstitutionCode: financialInstitutionCode,
		PaymentDescription:       paymentDescription,
	}
}

type PSEPayment struct {
	Payment
	PSEProduct
}

func NewPSEPayment(payment Payment, pseProduct PSEProduct) *PSEPayment {
	return &PSEPayment{Payment: payment, PSEProduct: pseProduct}
}
