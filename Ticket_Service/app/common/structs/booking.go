package commonStructs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RequestBookingControllerPayload struct {
	BookingId uuid.UUID `json:"bookingId" form:"bookingId" validate:"required"`
}

type RequestBookingServicePayload struct {
	Token     string    `json:"token" form:"token" validate:"required"`
	BookingId uuid.UUID `json:"bookingId" form:"bookingId" validate:"required"`
	TicketId  uuid.UUID `json:"ticketId" form:"ticketId" validate:"required"`
	UserId    string    `json:"userId" form:"userId" validate:"required"`
}

type RequestBookingServiceToken struct {
	jwt.RegisteredClaims
	UserId string `json:"userId"`
}
