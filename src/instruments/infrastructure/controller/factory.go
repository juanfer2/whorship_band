package instrument_controller

import (
	instrument_application "github.com/juanfer2/whorship_band/src/instruments/application"
	instrument_repositories "github.com/juanfer2/whorship_band/src/instruments/infrastructure/repository"
)

func InstrumentControllerFactory() *InstrumentController {
	instrumentRepository := instrument_repositories.NewInstrumentPGRepository()
	instrumentUseCase := instrument_application.NewInstrumentUseCase(instrumentRepository)

	return NewInstrumentController(instrumentUseCase)
}
