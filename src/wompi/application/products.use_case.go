package wompi_application

import (
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type ProductUseCase struct {
	productRepository wompi_repository.ProductRepository
}

func NewProductUseCase(productRepository wompi_repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepository: productRepository}
}

func (productUseCase *ProductUseCase) All() []wompi_entities.Product {
	return productUseCase.productRepository.All()
}

func (productUseCase *ProductUseCase) Create() []wompi_entities.Product {
	return productUseCase.productRepository.All()
}
