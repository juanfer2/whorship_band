package group_mate_application

import (
	group_mate_domain "github.com/juanfer2/whorship_band/src/group_mates/domain"
)

type GroupMateUseCase struct {
	repository group_mate_domain.GroupMateRepository
}

func NewGroupMateUseCase(
	repository group_mate_domain.GroupMateRepository,
) *GroupMateUseCase {
	return &GroupMateUseCase{repository: repository}
}

func (uc *GroupMateUseCase) FindByUuid(id string) []group_mate_domain.GroupMate {
	return uc.repository.FindByUuidWithInstruments(id)
}

func (uc *GroupMateUseCase) Create(
	groupMate group_mate_domain.GroupMate,
) (group_mate_domain.GroupMate, error) {
	newGroupMate, err := uc.repository.Create(groupMate)

	return newGroupMate, err
}

func (uc *GroupMateUseCase) CreateInBatches(groupMates []group_mate_domain.GroupMate) {
	uc.repository.CreateInBatches(groupMates)
}

func (uc *GroupMateUseCase) GetAll() []group_mate_domain.GroupMate {
	return uc.repository.All()
}

func (uc *GroupMateUseCase) Delete(id string) (group_mate_domain.GroupMate, error) {
	groupMate, err := uc.repository.FindByUuid(id)

	if err != nil {
		return groupMate, err
	}

	uc.repository.Delete(groupMate)

	return groupMate, nil
}

func (uc *GroupMateUseCase) Update(
	id string, groupMate group_mate_domain.GroupMate,
) (group_mate_domain.GroupMate, error) {
	updateGm, err := uc.repository.UpdateByUuid(id, groupMate)

	if err != nil {
		return groupMate, err
	}

	return updateGm, nil
}
