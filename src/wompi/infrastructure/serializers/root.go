package wompi_serializers

import wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"

type ProductSerializer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Bin      string `json:"bin"`
	Type     string `json:"type"`
	Token    string `json:"token"`
	SourceID int    `json:"source_id"`
	Mask     string `json:"mask"`
}

func NewProductSerializer(product wompi_entities.Product) *ProductSerializer {
	return &ProductSerializer{
		ID:       product.ID,
		Name:     product.Name,
		Bin:      product.Bin,
		Type:     product.Type.String(),
		Token:    product.Token,
		SourceID: product.SourceID,
		Mask:     product.Mask,
	}
}

func NewProductsSerializer(products []wompi_entities.Product) *[]ProductSerializer {
	var productsSerializer []ProductSerializer
	for _, product := range products {
		productsSerializer = append(productsSerializer, ProductSerializer{
			ID:       product.ID,
			Name:     product.Name,
			Bin:      product.Bin,
			Type:     product.Type.String(),
			Token:    product.Token,
			SourceID: product.SourceID,
			Mask:     product.Mask,
		})
	}

	return &productsSerializer
}
