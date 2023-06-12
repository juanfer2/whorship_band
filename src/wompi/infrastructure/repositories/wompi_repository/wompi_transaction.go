package wompi_repositories

import (
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
)

func (w *WompiRepository) ExecuteTransactionSourceCreditCard(
	paymentCreditCard wompi_entities.CreditCardSourcePayment,
) any {
	payment := PaymentWithSourceCreditCard(&paymentCreditCard)
	params := NewPaymentParams(payment)
	response := w.wompiClient.CreateTransactionWithCreditCard(*params)

	return response
}

func (w *WompiRepository) ExecuteTransactionPse(
	psePayment wompi_entities.PSEPayment,
) any {
	payment := PaymentWithPSE(&psePayment)
	params := NewPaymentParams(payment)

	return w.wompiClient.CreateTransaction(*params)
}
