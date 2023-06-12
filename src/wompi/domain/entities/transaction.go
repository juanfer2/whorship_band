package wompi_entities

import "time"

type StatusTransaction uint

const (
	PENDING StatusTransaction = iota
	APPROVED
	DECLINED
	ERROR
)

func (t StatusTransaction) String() string {
	switch t {
	case PENDING:
		return "PENDING"
	case APPROVED:
		return "APPROVED"
	case DECLINED:
		return "DECLINED"
	case ERROR:
		return "ERROR"
	}
	return "unknown"
}

type Transaction struct {
	ID             int               `json:"id"`
	SourceID       int               `json:"source_id"`
	Price          int               `json:"price"`
	Currency       string            `json:"currency"`
	CustomerData   CustomerData      `json:"customer_data"`
	Customer_email string            `json:"customer_email"`
	Reference      string            `json:"reference"`
	Status         StatusTransaction `json:"status"`

	///ProductID int     `json:"product_id"`
	//Product   Product `json:"product"`

	FinalizedAt time.Time `json:"finalized_at"`

	CreatedAt time.Time // `json:"created_at,omitempty"`
	UpdatedAt time.Time // `json:"updated_at,omitempty"`
}
