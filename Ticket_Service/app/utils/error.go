package utils

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorMessages struct {
	Errors any `json:"errors"`
}

type CustomError struct {
	StatusCode int   `json:"statusCode"`
	Err        error `json:"err"`
}

func (r *CustomError) Error() string {
	return r.Err.Error()
}
