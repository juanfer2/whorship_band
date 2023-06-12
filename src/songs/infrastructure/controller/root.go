package songs_controller

import (
	"github.com/gofiber/fiber/v2"

	middlewares "github.com/juanfer2/whorship_band/src/shared/infrastructure/controllers"
	songs_application "github.com/juanfer2/whorship_band/src/songs/application"
	songs_serializers "github.com/juanfer2/whorship_band/src/songs/infrastructure/serializers"
)

type SongController struct {
	songUseCase *songs_application.SongUseCase
}

func NewSongController(songUseCase *songs_application.SongUseCase) *SongController {
	return &SongController{songUseCase: songUseCase}
}

func (songController *SongController) All(c *fiber.Ctx) error {
	songs := songController.songUseCase.GetAll()

	return c.JSON(songs_serializers.NewSongSerializerSongList(songs))
}

func (songController *SongController) Create(c *fiber.Ctx) error {
	songParams, err := getParamsSong(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	song, err := songController.songUseCase.Create(songParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(songs_serializers.NewSongSerializerFromtSong(song))
}

func (songController *SongController) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	song, err := songController.songUseCase.FindByUuid(id)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(songs_serializers.NewSongSerializerFromtSong(song))
}

func (songController *SongController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	songParams, err := getParamsSong(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	song, err := songController.songUseCase.Update(id, songParams)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(songs_serializers.NewSongSerializerFromtSong(song))
}

func (songController *SongController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	song, err := songController.songUseCase.Delete(id)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(songs_serializers.NewSongSerializerFromtSong(song))
}
