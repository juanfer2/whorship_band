package wompi_grpc

import (
	wompipb "github.com/juanfer2/whorship_band/src/wompi/infrastructure/grpc/proto"
	"google.golang.org/grpc"

	wompi_application "github.com/juanfer2/whorship_band/src/wompi/application"
	product_repository "github.com/juanfer2/whorship_band/src/wompi/infrastructure/repositories/product_repository"
	wompi_repositories "github.com/juanfer2/whorship_band/src/wompi/infrastructure/repositories/wompi_repository"
)

func ServerWompiFactory() *Server {
	wompiRepository := wompi_repositories.NewWompiRepository()
	productRepository := product_repository.NewProductRepositoryPostgrest()
	tokenizeProduct := wompi_application.NewTokenizeProduct(wompiRepository, productRepository)

	return NewWompiServer(tokenizeProduct)
}

func RegisterServices(s *grpc.Server) {
	wompipb.RegisterWompiServiceServer(s, ServerWompiFactory())
}
