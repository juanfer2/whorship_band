package songs_serializers

import (
	"time"

	"github.com/google/uuid"
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
)

type SongSerializer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Tone      string    `json:"tone"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewSongSerializerFromtSong(song songs_domain.Song) *SongSerializer {
	return &SongSerializer{
		ID:        song.ID,
		Name:      song.Name,
		Url:       song.Url,
		Tone:      song.Url,
		CreatedAt: song.CreatedAt,
	}
}

func NewSongSerializerSongList(songs []songs_domain.Song) []SongSerializer {
	serializerList := make([]SongSerializer, 0, len(songs))

	for _, song := range songs {
		serializerList = append(serializerList, *NewSongSerializerFromtSong(song))
	}

	return serializerList
}
