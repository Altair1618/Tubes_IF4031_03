package models

import (
	"time"

	"github.com/google/uuid"
)

type TicketStatus string

const (
	Open    TicketStatus = "OPEN"
	Ongoing TicketStatus = "ONGOING"
	Booked  TicketStatus = "BOOKED"
)

type Ticket struct {
	Id        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventId   uuid.UUID    `gorm:"type:uuid;not null"`
	SeatId    string       `gorm:"not null"`
	Status    TicketStatus `gorm:"not null;default:OPEN"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
}
