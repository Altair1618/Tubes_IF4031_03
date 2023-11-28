package commonStructs

type JWTPayload struct {
	UserId    string
	SessionId string
	Secret    string
}
