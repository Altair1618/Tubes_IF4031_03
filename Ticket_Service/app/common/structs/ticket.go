package commonStructs

type TicketUpdateStatusRequest struct {
	InvoiceId string        `json:"invoiceId" form:"invoiceId" validate:"required"`
	Status    PaymentStatus `json:"status" form:"status" validate:"required,is_payment_status"`
}

type TicketUpdateStatusServicePayload struct {
	InvoiceId string
	Status    PaymentStatus
	UserId    string
}
