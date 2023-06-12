package wompi_repositories

import (
	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
)

type PaymentParamsOptions func(*wompi_client.PaymentBody)

func NewPaymentParams(paymentParamsOptions PaymentParamsOptions) *wompi_client.PaymentBody {
	var params wompi_client.PaymentBody
	paymentParamsOptions(&params)

	return &params
}

func PaymentWithSourceCreditCard(
	creditCardSource *wompi_entities.CreditCardSourcePayment,
) PaymentParamsOptions {
	return func(p *wompi_client.PaymentBody) {
		p.AmountInCents = creditCardSource.AmountInCents
		p.Currency = creditCardSource.Currency
		p.CustomerEmail = creditCardSource.CustomerEmail
		p.Reference = creditCardSource.Reference
		p.CustomerData = wompi_client.CustomerData(creditCardSource.CustomerData)
		p.PaymentMethod = wompi_client.PaymentMethodCreditCard{
			Type:         wompi_client.CREDIT_CARD.String(),
			Installments: creditCardSource.Installments,
			Token:        creditCardSource.Token,
		}
	}
}

func PaymentWithPSE(
	paymentParams *wompi_entities.PSEPayment,
) PaymentParamsOptions {
	return func(p *wompi_client.PaymentBody) {
		p.AmountInCents = paymentParams.AmountInCents
		p.Currency = paymentParams.Currency
		p.CustomerEmail = paymentParams.CustomerEmail
		p.Reference = paymentParams.Reference
		p.CustomerData = wompi_client.CustomerData(paymentParams.CustomerData)
		p.PaymentMethod = wompi_client.PaymentMethodPSE{
			Type:                     wompi_client.PSE.String(),
			UserType:                 paymentParams.UserType,
			UserLegalID:              paymentParams.UserLegalID,
			UserLegalIDType:          paymentParams.UserLegalIDType,
			FinancialInstitutionCode: paymentParams.FinancialInstitutionCode,
			PaymentDescription:       paymentParams.PaymentDescription,
		}
	}
}

func PaymentWithNequi(paymentMethodPSE wompi_entities.Nequi) PaymentParamsOptions {
	return func(p *wompi_client.PaymentBody) {
		p.PaymentMethod = paymentMethodPSE
	}
}
