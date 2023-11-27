package models

import (
	"time"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Ticket_Service/app/common/structs"
	"github.com/google/uuid"
)

type Ticket struct {
	Id        uuid.UUID                  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Price     int                        `gorm:"not null;column:price"`
	EventId   uuid.UUID                  `gorm:"type:uuid;not null;column:event_id;uniqueIndex:unique_seat_constraint"`
	SeatId    string                     `gorm:"not null;column:seat_id;uniqueIndex:unique_seat_constraint"`
	Status    commonStructs.TicketStatus `gorm:"not null;default:OPEN;column:status"`
	CreatedAt time.Time                  `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time                  `gorm:"autoUpdateTime;column:updated_at"`
}
