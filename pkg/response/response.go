package response

type response struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func BuildSuccess(code int, message string, data any) *response {
	return &response{
		Code:    code,
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func BuildError(code int, errorMessage string) *response {
	return &response{
		Code:    code,
		Error:   true,
		Message: errorMessage,
	}
}
