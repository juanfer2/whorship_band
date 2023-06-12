package wompi_application

import (
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type TokenizeProduct struct {
	wompiRepository   wompi_repository.WompiRepository
	productRepository wompi_repository.ProductRepository
}

func NewTokenizeProduct(
	wompiRepository wompi_repository.WompiRepository,
	productRepository wompi_repository.ProductRepository,
) *TokenizeProduct {
	return &TokenizeProduct{
		wompiRepository:   wompiRepository,
		productRepository: productRepository,
	}
}

func (tokenizeProduct *TokenizeProduct) TokenizeCreditCard(
	creditCard wompi_entities.CreditCard,
	customerEmail string,
) wompi_entities.Product {
	product := tokenizeProduct.wompiRepository.TokenizeCreditCard(creditCard, customerEmail)
	newProduct := tokenizeProduct.createProduct(product)

	return newProduct
}

func (tokenizeProduct *TokenizeProduct) createProduct(
	product wompi_entities.Product,
) wompi_entities.Product {
	newProduct, _ := tokenizeProduct.productRepository.Create(product)

	return newProduct
}
