package models

import (
	"time"

	commonStructs "github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/common/structs"
	"github.com/google/uuid"
)

type Invoice struct {
	Id           uuid.UUID                   `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Status       commonStructs.InvoiceStatus `gorm:"not null;default:ONGOING;column:status"`
	PaymentToken string                      `gorm:"not null;column:payment_token"`
	TicketId     uuid.UUID                   `gorm:"type:uuid;not null;column:ticket_id"`
	UserId       uuid.UUID                   `gorm:"type:uuid;not null;column:user_id"`
	CreatedAt    time.Time                   `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time                   `gorm:"autoUpdateTime;column:updated_at"`
}
