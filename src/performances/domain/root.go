package performances_domain

import (
	"time"

	"github.com/google/uuid"
	group_mate_domain "github.com/juanfer2/whorship_band/src/group_mates/domain"
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
	"gorm.io/gorm"
)

type Performance struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Date time.Time `json:"date"`

	PerformanceGroupMates []PerformanceGroupMate `json:"performanceGroupMate"`
	PerformanceSongs      []PerformanceSong      `json:"performanceSongs"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PerformanceSong struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	PerformanceID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"performanceId"`
	SongID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"songId"`

	Performance Performance       `json:"performance"`
	Song        songs_domain.Song `json:"song"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type PerformanceGroupMate struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	PerformanceID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"performanceId"`
	GroupMateID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"groutMateId"`

	Performance Performance                 `json:"performance"`
	GroupMate   group_mate_domain.GroupMate `json:"groutMate"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
export interface Event {
  id: string;
  date: Date;
  songs: Song[];
  members: Member[];
}

export interface Member {
  name: string;
  rol: string[];
  url?: string;
}

export interface EventForm {
  date: Date;
  songs?: Song[];
  members?: Member[];
}

export const EmptyEvent: EventForm = {
  date: new Date(),
  members: [],
  songs: []
};

export const EmptyEventObject: Event = {
  id: '',
  date: new Date(),
  members: [],
  songs: []
};

*/
