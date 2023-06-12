package performance_repositories

import (
	performances_domain "github.com/juanfer2/whorship_band/src/performances/domain"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

type PerformancePGRepository struct {
	postgres.PostgresRepository[performances_domain.Performance]
}

func NewPerformancePGRepository() performances_domain.PerformanceRepository {
	repository := PerformancePGRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}

func (p *PerformancePGRepository) AllWithSongsAndGroupMate() ([]performances_domain.Performance, error) {
	var performances []performances_domain.Performance
	err := p.PostgresRepository.Preload(
		"PerformanceGroupMates", "PerformanceGroupMates.GroupMate",
		"PerformanceSongs", "PerformanceSongs.Song",
	).Find(&performances).Error
	if err != nil {
		return performances, err
	}

	utilities.PrintJson(performances)

	return performances, nil
}

func (p *PerformancePGRepository) FindByUuidWithSongsAndGroupMates(
	id string,
) (performances_domain.Performance, error) {
	var performance performances_domain.Performance
	err := p.PostgresRepository.Preload(
		"PerformanceGroupMates", "PerformanceGroupMates.GroupMate",
		"PerformanceSongs", "PerformanceSongs.Song",
	).First(&performance).Error
	if err != nil {
		return performance, err
	}

	utilities.PrintJson(performance)

	return performance, nil
}
