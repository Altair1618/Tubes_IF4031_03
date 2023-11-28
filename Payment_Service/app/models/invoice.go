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
	Id         uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey:column:id"`
	Status     InvoiceStatus `gorm:"not null;default:ONGOING;column:status"`
	PaymentUrl string        `gorm:"not null"`
	CreatedAt  time.Time     `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt  time.Time     `gorm:"autoUpdateTime;column:updated_at"`
}
