package models

import (
	"time"

	"github.com/google/uuid"
)

type InvoiceStatus string

const (
	Success InvoiceStatus = "SUCCESS"
	Ongoing InvoiceStatus = "ONGOING"
	Failed  InvoiceStatus = "FAILED"
)

type Invoice struct {
	Id           uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Status       InvoiceStatus `gorm:"not null;default:ONGOING;column:status"`
	PaymentToken string        `gorm:"not null;column:payment_token"`
	TicketId     uuid.UUID     `gorm:"type:uuid;not null;column:ticket_id"`
	UserId       string        `gorm:"not null;column:user_id"`
	CreatedAt    time.Time     `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime;column:updated_at"`
}
