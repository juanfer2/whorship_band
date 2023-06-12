package instrument_repositories

import (
	intrument_domain "github.com/juanfer2/whorship_band/src/instruments/domain"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
)

type InstrumentPGRepository struct {
	postgres.PostgresRepository[intrument_domain.Instrument]
}

func NewInstrumentPGRepository() intrument_domain.InstrumentRepository {
	repository := InstrumentPGRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}
