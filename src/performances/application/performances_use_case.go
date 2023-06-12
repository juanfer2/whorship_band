package performances_application

import (
	performances_domain "github.com/juanfer2/whorship_band/src/performances/domain"
)

type PerformanceUseCase struct {
	repository performances_domain.PerformanceRepository
}

func NewPerformanceUseCase(repository performances_domain.PerformanceRepository) *PerformanceUseCase {
	return &PerformanceUseCase{repository: repository}
}

func (uc *PerformanceUseCase) FindById(id string) (performances_domain.Performance, error) {
	return uc.repository.FindByUuidWithSongsAndGroupMates(id)
}

func (
	uc *PerformanceUseCase,
) Create(instrument performances_domain.Performance) (performances_domain.Performance, error) {
	newInstrument, err := uc.repository.Create(instrument)
	return newInstrument, err
}

func (uc *PerformanceUseCase) CreateInBatches(instruments []performances_domain.Performance) {
	uc.repository.CreateInBatches(instruments)
}

func (uc *PerformanceUseCase) GetAll() ([]performances_domain.Performance, error) {
	return uc.repository.AllWithSongsAndGroupMate()
}

func (uc *PerformanceUseCase) Delete(id string) (performances_domain.Performance, error) {
	instrument, err := uc.repository.FindByUuid(id)
	if err != nil {
		return instrument, err
	}

	uc.repository.Delete(instrument)

	return instrument, nil
}

func (uc *PerformanceUseCase) Update(
	id string, instrument performances_domain.Performance,
) (performances_domain.Performance, error) {
	updateInst, err := uc.repository.UpdateByUuid(id, instrument)
	if err != nil {
		return instrument, err
	}

	return updateInst, nil
}
