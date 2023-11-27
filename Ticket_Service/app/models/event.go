package models

import (
	"time"

	"github.com/google/uuid"
)

// test

type Event struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EventName string    `gorm:"not null" json:"event_name"`
	EventTime time.Time `gorm:"not null" json:"event_time"`
	Location  string    `gorm:"not null" json:"location"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
