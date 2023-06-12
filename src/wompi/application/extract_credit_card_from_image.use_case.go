package wompi_application

import (
	"github.com/juanfer2/whorship_band/src/shared/domain"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type ExtractCreditcCardFromImage struct {
	readImageRepository wompi_repository.ReadImageRepository
}

func NewExtractCreditcCardFromImage(
	readImageRepository wompi_repository.ReadImageRepository,
) *ExtractCreditcCardFromImage {
	return &ExtractCreditcCardFromImage{readImageRepository: readImageRepository}
}

func (e ExtractCreditcCardFromImage) GetCreditCard(file domain.File) wompi_entities.CreditCard {
	return e.readImageRepository.ExtractText(file)
}
