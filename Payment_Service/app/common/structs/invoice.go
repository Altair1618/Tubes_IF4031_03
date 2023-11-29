package commonStructs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type InvoiceStatus string

const (
	Success InvoiceStatus = "SUCCESS"
	Ongoing InvoiceStatus = "ONGOING"
	Failed  InvoiceStatus = "FAILED"
)

type CreateInvoiceControllerPayload struct {
	TicketId uuid.UUID `json:"ticketId" form:"ticketId" validate:"required"`
}

type CreateInvoiceServicePayload struct {
	TicketId uuid.UUID `json:"ticketId" form:"ticketId" validate:"required"`
	UserId   string    `json:"userId" form:"userId" validate:"required"`
}

type InvoiceTokenClaims struct {
	jwt.RegisteredClaims
	TicketId uuid.UUID
	UserId   string
}
