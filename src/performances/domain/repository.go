package performances_domain

type PerformanceRepository interface {
	FindBy(id int) Performance
	FindByUuid(id string) (Performance, error)
	FindByUuidWithSongsAndGroupMates(id string) (Performance, error)
	Create(performance Performance) (Performance, error)
	Update(performance Performance) (Performance, error)
	UpdateByUuid(id string, updateRecord Performance) (Performance, error)
	CreateInBatches(performances []Performance)
	All() []Performance
	AllWithSongsAndGroupMate() ([]Performance, error)
	Delete(performance Performance) Performance
	WhereBy(query interface{}, arg ...interface{}) Performance
}
