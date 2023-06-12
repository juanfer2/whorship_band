package wompi_application

import (
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type ExecuteTransaction struct {
	wompiRepository wompi_repository.WompiRepository
}

func NewExecuteTransaction(wompiRepository wompi_repository.WompiRepository) *ExecuteTransaction {
	return &ExecuteTransaction{wompiRepository: wompiRepository}
}

func (
	executeTransaction *ExecuteTransaction,
) PayWithSourceCreditCard(paymentInfo wompi_entities.CreditCardSourcePayment) any {
	return executeTransaction.wompiRepository.ExecuteTransactionSourceCreditCard(paymentInfo)
}

func (
	executeTransaction *ExecuteTransaction,
) PayWithPSE(paymentInfo wompi_entities.PSEPayment) any {
	return executeTransaction.wompiRepository.ExecuteTransactionPse(paymentInfo)
}
