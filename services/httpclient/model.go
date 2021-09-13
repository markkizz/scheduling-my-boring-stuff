package httpclient

import "net/http"

type BaseHttpClient interface {
	request(req *http.Request, v interface{}) error
}

type ErrorResponse struct {
	// Code    int    `json:"code"`
	// Message string `json:"message"`
	Message string `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
