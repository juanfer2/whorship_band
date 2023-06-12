package songs_domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `json:"name"`
	Url  string    `json:"url"`
	Tone string    `json:"tone"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
