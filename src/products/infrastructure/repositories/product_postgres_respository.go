package products_repositories

import (
	products_domain "github.com/juanfer2/whorship_band/src/products/domain"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
)

type ProductPostgresRepository struct {
	postgres.PostgresRepository[products_domain.Product]
}

func NewProductPostgresRepository() products_domain.ProductRepository {
	repository := ProductPostgresRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}

func (p *ProductPostgresRepository) Where(query any, arg ...any) (products []products_domain.Product) {
	return
}
