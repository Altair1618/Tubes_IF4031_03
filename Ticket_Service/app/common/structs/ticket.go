package commonStructs

import "github.com/google/uuid"

type TicketStatus string

const (
	Open    TicketStatus = "OPEN"
	Ongoing TicketStatus = "ONGOING"
	Booked  TicketStatus = "BOOKED"
)

type CreateTicketServicePayload struct {
	Price   int       `json:"price" form:"price" validate:"required"`
	EventId uuid.UUID `json:"eventId" form:"eventId" validate:"required"`
	SeatId  string    `json:"seatId" form:"seatId" validate:"required,is_seat_number"`
}

type UpdateTicketServicePayload struct {
	Price   int          `json:"price" form:"price"`
	EventId uuid.UUID    `json:"eventId" form:"eventId"`
	SeatId  string       `json:"seatId" form:"seatId" validate:"is_seat_number"`
	Status  TicketStatus `json:"status" form:"status"`
}

type UpdateStatusServicePayload struct {
	InvoiceId string        `json:"invoiceId" form:"invoiceId"`
	Status    PaymentStatus `json:"status" form:"status"`
	UserId    string        `json:"userId" form:"userId"`
}
