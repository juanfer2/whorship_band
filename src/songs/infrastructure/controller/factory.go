package songs_controller

import (
	songs_application "github.com/juanfer2/whorship_band/src/songs/application"
	songs_repositories "github.com/juanfer2/whorship_band/src/songs/infrastructure/repositories"
)

func SongControllerFactory() *SongController {
	songRepository := songs_repositories.NewSongPGRepository()
	songUseCase := songs_application.NewSongUseCase(songRepository)

	return NewSongController(songUseCase)
}
