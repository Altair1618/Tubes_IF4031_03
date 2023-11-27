package models

import "github.com/google/uuid"

type TicketInvoiceBooking struct {
	Id        uint      `gorm:"primaryKey;autoIncrement;column:id"`
	InvoiceId uuid.UUID `gorm:"type:uuid;not null;index;column:invoice_id"`
	TicketId  uuid.UUID `gorm:"type:uuid;not null;column:ticket_id"`
	BookingId uuid.UUID `gorm:"type:uuid;not null;column:booking_id"`
}
