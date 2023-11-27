package models

import (
	"time"

	"github.com/google/uuid"
)

// test

type Event struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column=id"`
	EventName string    `gorm:"not null;column=event_name"`
	EventTime time.Time `gorm:"not null;column=event_time"`
	Location  string    `gorm:"not null;column=event_location"`
	CreatedAt time.Time `gorm:"autoCreateTime;column=created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column=updated_at"`
}
