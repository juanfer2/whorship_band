package wompi_repository

import (
	"github.com/juanfer2/whorship_band/src/shared/domain"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
)

type WompiRepository interface {
	ExecuteTransactionPse(psePayment wompi_entities.PSEPayment) any
	TokenizeNequi(creditCard wompi_entities.Nequi)
	GetMerchant() wompi_entities.Merchant
	ExecuteTransactionSourceCreditCard(
		paymentCreditCard wompi_entities.CreditCardSourcePayment,
	) any
	TokenizeCreditCard(
		creditCard wompi_entities.CreditCard, customerEmail string,
	) wompi_entities.Product
}

type DefaultRepository[T any] interface {
	FindBy(id int) T
	Create(newRecord T) (T, error)
	All() []T
	Delete(record T) T
	WhereBy(query any, arg ...any) T
	// FindWithAssociation(associations []string, query any, arg ...any) T
	// WhereWithAssociation(associations []string, query any, arg ...any) []T
	// Join(relations ...string) *gorm.DB
	CreateInBatches(records []T)
}

type ProductRepository interface {
	DefaultRepository[wompi_entities.Product]
}

type ReadImageRepository interface {
	ExtractText(file domain.File) wompi_entities.CreditCard
}
