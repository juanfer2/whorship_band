package group_mate_domain

type GroupMateRepository interface {
	FindBy(id int) GroupMate
	FindByUuid(id string) (GroupMate, error)
	Create(groupMate GroupMate) (GroupMate, error)
	Update(groupMate GroupMate) (GroupMate, error)
	UpdateByUuid(id string, updateRecord GroupMate) (GroupMate, error)
	CreateInBatches(groupMates []GroupMate)
	All() []GroupMate
	FindByUuidWithInstruments(id string) []GroupMate
	WhereWithAssociation() []GroupMate
	FindWithAssociation(associations []string, query any, arg ...any) GroupMate
	Delete(groupMate GroupMate) GroupMate
	WhereBy(query any, arg ...any) GroupMate
}
