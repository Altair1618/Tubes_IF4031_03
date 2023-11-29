package commonStructs

type ProcessPaymentServicePayload struct {
	PaymentToken string
	UserId       string
	JWTToken     string
}

type ProcessPaymentRequest struct {
	PaymentToken string `params:"paymentToken" validate:"required"`
}
