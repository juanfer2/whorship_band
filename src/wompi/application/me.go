package wompi_application

import wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"

type Me struct {
	wompiRepository wompi_repository.WompiRepository
}

func (m *Me) Info() any {
	return m.wompiRepository.GetMerchant()
}
