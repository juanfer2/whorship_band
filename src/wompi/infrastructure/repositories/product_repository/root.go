package product_repository

import (
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type ProductRepositoryPostgrest struct {
	postgres.PostgresRepository[wompi_entities.Product]
}

type ProductPostgresRepository struct {
	postgres.PostgresRepository[wompi_entities.Product]
}

func NewProductRepositoryPostgrest() wompi_repository.ProductRepository {
	repository := ProductPostgresRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}
