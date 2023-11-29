package commonStructs

type JWTPayload struct {
	UserId    string
	SessionId string
	Secret    string
}

type HttpResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
