package songs_domain

type SongRepository interface {
	FindBy(id int) Song
	FindByUuid(id string) (Song, error)
	Create(Song Song) (Song, error)
	Update(Song Song) (Song, error)
	UpdateByUuid(id string, updateRecord Song) (Song, error)
	CreateInBatches(Songs []Song)
	All() []Song
	WhereWithAssociation(associations []string, query any, arg ...any) []Song
	FindWithAssociation(associations []string, query any, arg ...any) Song
	Delete(Song Song) Song
	WhereBy(query any, arg ...any) Song
}
