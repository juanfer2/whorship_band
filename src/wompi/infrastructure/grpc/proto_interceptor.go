package wompi_grpc

import (
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompipb "github.com/juanfer2/whorship_band/src/wompi/infrastructure/grpc/proto"
)

func TransactionFromProto(
	req *wompipb.TransactionRequest,
) wompi_entities.Transaction {
	return wompi_entities.Transaction{
		SourceID:       int(*req.SourceId),
		Price:          int(req.Price),
		Currency:       req.Currency,
		Customer_email: req.Costumer.Email,
		Reference:      req.Reference,
		CustomerData: wompi_entities.CustomerData{
			PhoneNumber: req.Costumer.FullName,
			FullName:    req.Costumer.FullName,
		},
	}
}

func TransactionFromProtoInterceptor() wompi_entities.CreditCardSourcePayment {
	return wompi_entities.CreditCardSourcePayment{}
}
