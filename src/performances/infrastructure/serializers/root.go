package performances_serializer

import (
	"time"

	"github.com/google/uuid"
	group_mate_domain "github.com/juanfer2/whorship_band/src/group_mates/domain"
	group_mates_serializers "github.com/juanfer2/whorship_band/src/group_mates/infrastructure/serializers"
	performances_domain "github.com/juanfer2/whorship_band/src/performances/domain"
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
	songs_serializers "github.com/juanfer2/whorship_band/src/songs/infrastructure/serializers"
)

type PerformanceSerializer struct {
	ID   uuid.UUID `json:"id"`
	Date time.Time `json:"date"`

	Songs      []songs_serializers.SongSerializer            `json:"songs"`
	GroupMates []group_mates_serializers.GroupMateSerializer `json:"groupMates"`

	CreatedAt time.Time `json:"createdAt"`
}

func NewPerformanceSerializerFromItem(
	performance performances_domain.Performance,
) *PerformanceSerializer {
	// Songs
	var songs []songs_domain.Song

	for _, performanceSong := range performance.PerformanceSongs {
		songs = append(songs, performanceSong.Song)
	}

	songList := songs_serializers.NewSongSerializerSongList(songs)

	// groupMates
	var groupMates []group_mate_domain.GroupMate

	for _, performanceGroupMates := range performance.PerformanceGroupMates {
		groupMates = append(groupMates, performanceGroupMates.GroupMate)
	}

	groupMatesList := group_mates_serializers.NewGroupMateSerializerFromGroupMateList(groupMates)

	return &PerformanceSerializer{
		ID:         performance.ID,
		Date:       performance.Date,
		Songs:      songList,
		GroupMates: groupMatesList,
	}
}

func NewPerformanceSerializerInstrumentList(
	performances []performances_domain.Performance,
) []PerformanceSerializer {
	serializerList := make([]PerformanceSerializer, 0, len(performances))

	for _, performance := range performances {
		serializerList = append(serializerList, *NewPerformanceSerializerFromItem(performance))
	}

	return serializerList
}
