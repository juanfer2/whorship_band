package group_mate_repositories

import (
	group_mate_domain "github.com/juanfer2/whorship_band/src/group_mates/domain"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
)

type GroupMatePGRepository struct {
	postgres.PostgresRepository[group_mate_domain.GroupMate]
}

func NewGroupMatePGRepository() group_mate_domain.GroupMateRepository {
	repository := GroupMatePGRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}

func (p *GroupMatePGRepository) Where(query any, arg ...any) (GroupMate []group_mate_domain.GroupMate) {
	return
}

func (p *GroupMatePGRepository) FindByUuidWithInstruments(id string) []group_mate_domain.GroupMate {
	var groupMates []group_mate_domain.GroupMate
	p.PostgresRepository.Preload("Instruments", "Instruments.Instrument").Find(&groupMates)

	return groupMates
}

func (p *GroupMatePGRepository) WhereWithAssociation() []group_mate_domain.GroupMate {
	var groupMates []group_mate_domain.GroupMate
	p.PostgresRepository.Preload("Instruments", "Instruments.Instrument").Find(&groupMates)

	return groupMates
}
