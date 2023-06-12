package group_mate_domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupMate struct {
	gorm.Model
	// ID       string `json:"id" gorm:"serializer:json"`
	ID          uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string                `json:"name"`
	URLImage    string                `json:"urlImage"`
	Instruments []GroupMateInstrument `json:"instruments" `
	//Songs       []Song `json:"id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt time.Time
}

type Song struct {
	Name string
	Tone string
}
