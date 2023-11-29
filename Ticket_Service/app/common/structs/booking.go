package commonStructs

import "github.com/google/uuid"

type RequestBookingControllerPayload struct {
	BookingId uuid.UUID `json:"bookingId" form:"bookingId" validate:"required"`
}

type RequestBookingServicePayload struct {
	BookingId uuid.UUID `json:"bookingId" form:"bookingId" validate:"required"`
	TicketId  uuid.UUID `json:"ticketId" form:"ticketId" validate:"required"`
	UserId    string    `json:"userId" form:"userId" validate:"required"`
}
