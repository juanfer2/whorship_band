package songs_application

import (
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
)

type SongUseCase struct {
	repository songs_domain.SongRepository
}

func NewSongUseCase(repository songs_domain.SongRepository) *SongUseCase {
	return &SongUseCase{repository: repository}
}

func (songUseCase *SongUseCase) FindByUuid(id string) (songs_domain.Song, error) {
	return songUseCase.repository.FindByUuid(id)
}

func (songUseCase *SongUseCase) Create(song songs_domain.Song) (songs_domain.Song, error) {
	return songUseCase.repository.Create(song)
}

func (songUseCase *SongUseCase) CreateInBatches(songs []songs_domain.Song) {
	songUseCase.repository.CreateInBatches(songs)
}

func (songUseCase *SongUseCase) GetAll() []songs_domain.Song {
	return songUseCase.repository.All()
}

func (songUseCase *SongUseCase) Delete(id string) (songs_domain.Song, error) {
	song, err := songUseCase.repository.FindByUuid(id)
	if err != nil {
		return song, err
	}

	songUseCase.repository.Delete(song)

	return song, nil
}

func (songUseCase *SongUseCase) Update(
	id string, song songs_domain.Song,
) (songs_domain.Song, error) {
	updateSong, err := songUseCase.repository.UpdateByUuid(id, song)

	if err != nil {
		return song, err
	}

	return updateSong, nil
}
