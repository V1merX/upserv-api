package response

type Response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   bool        `json:"error"`
}

func BuildSuccess(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func BuildError(code int, errorMessage string) Response {
	return Response{
		Code:    code,
		Error:   true,
		Message: errorMessage,
	}
}
