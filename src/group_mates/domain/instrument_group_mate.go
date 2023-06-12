package group_mate_domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupMateInstrument struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	GroupMateID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"groupMateId"`
	InstrumentID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"instrumentId"`

	GroupMate  GroupMate  `json:"groupMate"`
	Instrument Instrument `json:"instrument"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
