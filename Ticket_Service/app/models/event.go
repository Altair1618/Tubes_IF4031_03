package models

import (
	"time"

	"github.com/google/uuid"
)

// test

type Event struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventName string    `gorm:"not null"`
	EventTime time.Time `gorm:"not null"`
	Location  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
