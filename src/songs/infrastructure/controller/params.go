package songs_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
)

type SongParams struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Tone string `json:"tone"`
}

func getParamsSong(c *fiber.Ctx) (songs_domain.Song, error) {
	var songParams SongParams
	if err := c.BodyParser(&songParams); err != nil {
		return songs_domain.Song{}, err
	}

	utilities.PrintJson(songParams)

	fmt.Println(songParams)

	return songs_domain.Song{
		Name: songParams.Name,
		Url:  songParams.Url,
		Tone: songParams.Tone,
	}, nil
}
