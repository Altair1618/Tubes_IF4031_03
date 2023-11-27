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
	Id        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EventId   uuid.UUID    `gorm:"type:uuid;not null" json:"event_id"`
	SeatId    string       `gorm:"not null" json:"seat_id"`
	Status    TicketStatus `gorm:"not null;default:OPEN" json:"status"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
