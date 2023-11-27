package commonStructs

type UpdateStatusServicePayload struct {
	InvoiceId string        `json:"invoiceId" form:"invoiceId"`
	Status    PaymentStatus `json:"status" form:"status"`
	UserId    string        `json:"userId" form:"userId"`
}
