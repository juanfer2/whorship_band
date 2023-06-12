package wompi_grpc

import (
	context "context"
	"fmt"
	"strconv"

	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompipb "github.com/juanfer2/whorship_band/src/wompi/infrastructure/grpc/proto"
)

func (s *Server) CreateCreditCardProduct(
	ctx context.Context, req *wompipb.CreditCardParams,
) (*wompipb.CreditCardProductMessage, error) {
	product := s.tokenizeProduct.TokenizeCreditCard(wompi_entities.CreditCard{
		Number:     req.Number,
		ExpMonth:   req.ExpMonth,
		ExpYear:    req.ExpYear,
		CVC:        req.Cvc,
		CardHolder: req.CardHolder,
	}, req.CustomerEmail)

	productMessage := &wompipb.CreditCardProductMessage{
		Id:       strconv.Itoa(product.ID),
		Name:     product.Name,
		Bin:      product.Bin,
		Type:     product.Type.String(),
		SourceID: int32(product.SourceID),
		Token:    product.Token,
		Mask:     product.Mask,
		//CreatedAt: product.CreatedAt,
		//UpdatedAt: product.UpdatedAt,
	}

	fmt.Println(product)

	return productMessage, nil
}
