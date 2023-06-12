package wompi_grpc

import (
	wompi_application "github.com/juanfer2/whorship_band/src/wompi/application"
	wompipb "github.com/juanfer2/whorship_band/src/wompi/infrastructure/grpc/proto"
)

type Server struct {
	wompipb.UnimplementedWompiServiceServer
	tokenizeProduct *wompi_application.TokenizeProduct
}

func NewWompiServer(
	tokenizeProduct *wompi_application.TokenizeProduct,
) *Server {
	return &Server{tokenizeProduct: tokenizeProduct}
}
