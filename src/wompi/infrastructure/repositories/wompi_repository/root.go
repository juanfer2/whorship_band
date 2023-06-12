package wompi_repositories

import (
	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	wompi_products "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi/products"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type WompiRepository struct {
	wompiClient *wompi_client.WompiClient
}

func NewWompiRepository() wompi_repository.WompiRepository {
	wompiClient := wompi_client.WompiClientFactory()
	return &WompiRepository{wompiClient: wompiClient}
}

func (w *WompiRepository) TokenizeNequi(nequi wompi_entities.Nequi) {
	productNequi := wompi_client.TokenizeNequiOption(
		wompi_products.Nequi{PhoneNumber: nequi.PhoneNumber},
	)

	w.wompiClient.Tokenize(wompi_client.NewTokenizeParams(productNequi))
}

func (w *WompiRepository) GetMerchant() wompi_entities.Merchant {
	merchant := w.wompiClient.GetMerchant()
	return wompi_entities.Merchant{
		ID:                     merchant.ID,
		Name:                   merchant.Name,
		Email:                  merchant.Email,
		ContactName:            merchant.ContactName,
		PhoneNumber:            merchant.ContactName,
		Active:                 merchant.Active,
		LogoURL:                merchant.LogoURL,
		LegalName:              merchant.LegalName,
		LegalIDType:            merchant.LegalIDType,
		LegalID:                merchant.LegalID,
		PublicKey:              merchant.PublicKey,
		AcceptedCurrencies:     merchant.AcceptedCurrencies,
		FraudJavascriptKey:     merchant.FraudJavascriptKey,
		FraudGroups:            merchant.FraudGroups,
		AcceptedPaymentMethods: merchant.AcceptedPaymentMethods,
		PaymentMethods:         ConvertPaymentMethodsjson(merchant.PaymentMethods),
	}
}

func ConvertPaymentMethodsjson(a []wompi_client.PaymentMethod) (b []wompi_entities.PaymentMethod) {
	for _, v := range a {
		r, _ := utilities.TypeConverter[wompi_entities.PaymentMethod](v)
		b = append(b, *r)
	}

	return
}
