package wompi_repositories

import (
	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
)

func convertPaymentMethods(
	paymentsMethods []wompi_client.PaymentMethod,
) (newPaymentMethods []wompi_entities.PaymentMethod) {

	for _, paymentsMethod := range paymentsMethods {
		newPaymentMethods = append(newPaymentMethods, wompi_entities.PaymentMethod{
			Name:              paymentsMethod.Name,
			PaymentProcessors: convertPaymentProcessors(paymentsMethod.PaymentProcessors),
		})
	}

	return
}

func convertPaymentProcessors(
	paymentProcessors []wompi_client.PaymentProcessor,
) (newPaymentProcessors []wompi_entities.PaymentProcessor) {
	for _, paymentProcessors := range paymentProcessors {
		newPaymentProcessors = append(newPaymentProcessors, wompi_entities.PaymentProcessor{
			Name: paymentProcessors.Name,
		})
	}

	return
}
