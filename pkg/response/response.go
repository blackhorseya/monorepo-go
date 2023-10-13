package response

import (
	"net/http"
)

var (
	// OK request is successful.
	OK = &Response{Code: http.StatusOK, Message: "ok"}

	// Err request is failed.
	Err = &Response{Code: http.StatusInternalServerError, Message: "error"}
)

// Response response struct.
type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WithData set response data.
func (resp *Response) WithData(data any) *Response {
	return &Response{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    data,
	}
}
