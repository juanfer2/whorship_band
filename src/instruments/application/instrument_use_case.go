package instrument_application

import (
	intrument_domain "github.com/juanfer2/whorship_band/src/instruments/domain"
)

type InstrumentUseCase struct {
	repository intrument_domain.InstrumentRepository
}

func NewInstrumentUseCase(repository intrument_domain.InstrumentRepository) *InstrumentUseCase {
	return &InstrumentUseCase{repository: repository}
}

func (uc *InstrumentUseCase) FindById(id int) intrument_domain.Instrument {
	return uc.repository.FindBy(id)
}

func (uc *InstrumentUseCase) Create(instrument intrument_domain.Instrument) (intrument_domain.Instrument, error) {
	newInstrument, err := uc.repository.Create(instrument)
	return newInstrument, err
}

func (uc *InstrumentUseCase) CreateInBatches(instruments []intrument_domain.Instrument) {
	uc.repository.CreateInBatches(instruments)
}

func (uc *InstrumentUseCase) GetAll() []intrument_domain.Instrument {
	return uc.repository.All()
}

func (uc *InstrumentUseCase) Delete(id string) (intrument_domain.Instrument, error) {
	instrument, err := uc.repository.FindByUuid(id)
	if err != nil {
		return instrument, err
	}

	uc.repository.Delete(instrument)

	return instrument, nil
}

func (uc *InstrumentUseCase) Update(id string, instrument intrument_domain.Instrument) (intrument_domain.Instrument, error) {
	updateInst, err := uc.repository.UpdateByUuid(id, instrument)
	if err != nil {
		return instrument, err
	}

	return updateInst, nil
}
