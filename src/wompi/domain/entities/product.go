package wompi_entities

import (
	"time"

	"gorm.io/gorm"
)

type TypeProduct uint

const (
	PRODUCT_CREDIT_CARD TypeProduct = iota

	PRODUCT_NEQUI
)

func (t TypeProduct) String() string {
	switch t {
	case PRODUCT_CREDIT_CARD:
		return "CARD"
	case PRODUCT_NEQUI:
		return "NEQUI"
	}
	return "unknown"
}

type Product struct {
	gorm.Model
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Bin      string      `json:"bin"`
	Type     TypeProduct `json:"type"`
	Token    string      `json:"token"`
	SourceID int         `json:"source"`
	Mask     string      `json:"mask"`

	CreatedAt time.Time // `json:"created_at,omitempty"`
	UpdatedAt time.Time // `json:"updated_at,omitempty"`
}

func (product *Product) IsNequi() bool {
	return product.Type == PRODUCT_NEQUI
}

func (product *Product) IsCreditCard() bool {
	return product.Type == PRODUCT_CREDIT_CARD
}

func (product *Product) TypeProduct() string {
	return product.Type.String()
}
