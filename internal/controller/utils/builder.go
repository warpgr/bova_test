package utils

// BuildOkResponse builds ok response.
func BuildOkResponse[DataT any](data DataT, message string) Response[DataT] {
	return Response[DataT]{
		Code:    0,
		Message: message,
		Data:    data,
	}
}

// BuildErrorResponse build error response.
func BuildErrorResponse(err error, code int) Response[interface{}] {
	return Response[interface{}]{
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	}
}
