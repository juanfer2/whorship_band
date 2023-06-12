package wompi_controller

import (
	wompi_application "github.com/juanfer2/whorship_band/src/wompi/application"
	product_repository "github.com/juanfer2/whorship_band/src/wompi/infrastructure/repositories/product_repository"
	"github.com/juanfer2/whorship_band/src/wompi/infrastructure/repositories/read_image_repository"
	wompi_repositories "github.com/juanfer2/whorship_band/src/wompi/infrastructure/repositories/wompi_repository"
)

func FactoryWompiController() *WompiController {
	wompiRepository := wompi_repositories.NewWompiRepository()
	productRepository := product_repository.NewProductRepositoryPostgrest()
	readImage := read_image_repository.NewReadImageRepository()

	tokenizeService := wompi_application.NewTokenizeProduct(wompiRepository, productRepository)
	productUseCase := wompi_application.NewProductUseCase(productRepository)
	transactionUseCase := wompi_application.NewExecuteTransaction(wompiRepository)

	extractCreditCardUseCase := wompi_application.NewExtractCreditcCardFromImage(readImage)

	return NewWompiController(tokenizeService, productUseCase, extractCreditCardUseCase,
		transactionUseCase)
}
