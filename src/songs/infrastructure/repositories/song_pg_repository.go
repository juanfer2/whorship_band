package songs_repositories

import (
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/persistence/postgres"
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
)

type SongPGRepository struct {
	postgres.PostgresRepository[songs_domain.Song]
}

func NewSongPGRepository() songs_domain.SongRepository {
	repository := SongPGRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}

func (p *SongPGRepository) Where(query any, arg ...any) (GroupMate []songs_domain.Song) {
	return
}
