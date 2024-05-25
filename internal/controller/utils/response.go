package utils

// Response common format of response body.
type Response[DataT any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    DataT  `json:"data"`
}
