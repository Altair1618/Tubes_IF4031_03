package commonStructs

type PaymentStatus string

const (
	Failed  PaymentStatus = "FAILED"
	Success PaymentStatus = "SUCCESS"
)

type SuccessPDFPayload struct {
	Price int
	Seat  string
}

type FailedPDFPayload struct {
	ErrorMessage string
}

type JWTPayload struct {
	UserId    string
	SessionId string
	Secret    string
}
