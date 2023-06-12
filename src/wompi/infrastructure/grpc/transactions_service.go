package wompi_grpc

import (
	context "context"

	wompipb "github.com/juanfer2/whorship_band/src/wompi/infrastructure/grpc/proto"
)

func (s *Server) CreateTransacction(
	ctx context.Context, req *wompipb.TransactionRequest,
) (*wompipb.TransactionMessage, error) {
	//params := TransactionFromProto(req)
	//s.executeTransaction.PayWithSourceCreditCard()

	return &wompipb.TransactionMessage{
		Reference:  req.Reference,
		CreditCard: req.CreditCard,
		SourceId:   req.SourceId,
		Price:      req.Price,
		Costumer:   req.Costumer,
	}, nil
}

func (s *Server) GetTransactionById(
	ctx context.Context, req *wompipb.GetTransactionParams,
) (*wompipb.TransactionMessage, error) {

	return &wompipb.TransactionMessage{}, nil
}
