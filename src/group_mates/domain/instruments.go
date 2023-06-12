package group_mate_domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Instrument struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `json:"name"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
