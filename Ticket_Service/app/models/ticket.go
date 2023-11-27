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
	Id        uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Price     int          `gorm:"not null;column:price"`
	EventId   uuid.UUID    `gorm:"type:uuid;not null;column:event_id"`
	SeatId    string       `gorm:"not null;column:seat_id"`
	Status    TicketStatus `gorm:"not null;default:OPEN;column:status"`
	CreatedAt time.Time    `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime;column:updated_at"`
}
