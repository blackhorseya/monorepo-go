package response

import (
	"context"
	"encoding/json"
	"net/http"
)

var (
	// OK request is successful.
	OK = &Response{Code: http.StatusOK, Message: "ok"}

	// Err request is failed.
	Err = &Response{Code: http.StatusInternalServerError, Message: "unknown error"}
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

// WrapError wrap error to response.
func (resp *Response) WrapError(err error) *Response {
	return &Response{
		Code:    resp.Code,
		Message: err.Error(),
		Data:    nil,
	}
}

// EncodeJSON encode response to json.
func EncodeJSON(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
