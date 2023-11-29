package commonStructs

type ProcessPaymentServicePayload struct {
	PaymentToken string
	UserId       string
}

type ProcessPaymentRequest struct {
	PaymentToken string `params:"paymentToken" validate:"required"`
}
